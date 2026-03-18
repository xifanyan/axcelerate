#!/usr/bin/env python3
"""
ADP Output Types Generator

Generates language-specific type definitions from outputs.json schema.

Usage:
    python scripts/generate_output_types.py --lang=go --output=src/go/
    python scripts/generate_output_types.py --lang=python --output=src/python/
    python scripts/generate_output_types.py --lang=ts --output=src/ts/
"""

import argparse
import json
import os
import sys
from pathlib import Path
from typing import Any, Dict, List, Optional


TYPE_MAPPINGS = {
    "go": {
        "string": "string",
        "integer": "int",
        "boolean": "bool",
        "array": "[]",
        "map_string_string_array": "map[string][]string",
    },
    "python": {
        "string": "str",
        "integer": "int",
        "boolean": "bool",
        "array": "List",
        "map_string_string_array": "Dict[str, List[str]]",
    },
    "ts": {
        "string": "string",
        "integer": "number",
        "boolean": "boolean",
        "array": "",
        "map_string_string_array": "Record<string, string[]>",
    },
    "rust": {
        "string": "String",
        "integer": "i32",
        "boolean": "bool",
        "array": "Vec",
        "map_string_string_array": "HashMap<String, Vec<String>>",
    },
}


def get_type(ref: str, definitions: Dict, lang: str) -> str:
    """Resolve $ref to actual type name."""
    if ref.startswith("#/definitions/"):
        type_name = ref.split("/")[-1]
        return type_name
    return ref


def get_field_type(field: Dict, definitions: Dict, lang: str) -> str:
    """Get the Go/Python/TypeScript/Rust type for a field."""
    # Handle $ref directly (just a type name string)
    if isinstance(field, str):
        return get_type(field, definitions, lang)

    field_type = field.get("type", "")

    if "$ref" in field:
        ref = field["$ref"]
        # Handle $ref as string (just the type name)
        if isinstance(ref, str):
            ref_type = get_type(ref, definitions, lang)
            if field_type == "array":
                if lang == "rust":
                    return f"Vec<{ref_type}>"
                return f"{TYPE_MAPPINGS[lang]['array']}{ref_type}"
            return ref_type
        # Handle $ref as dict
        ref_type = get_field_type(ref, definitions, lang)
        if field_type == "array":
            if lang == "rust":
                return f"Vec<{ref_type}>"
            return f"{TYPE_MAPPINGS[lang]['array']}{ref_type}"
        return ref_type

    if field_type == "array":
        items = field.get("items", {})
        if isinstance(items, str):
            # Just a type name
            item_type = get_type(items, definitions, lang)
            if lang == "rust":
                return f"Vec<{item_type}>"
            return f"{TYPE_MAPPINGS[lang]['array']}{item_type}"
        if "$ref" in items:
            item_type = get_field_type(items["$ref"], definitions, lang)
            if lang == "rust":
                return f"Vec<{item_type}>"
            return f"{TYPE_MAPPINGS[lang]['array']}{item_type}"
        base_type = TYPE_MAPPINGS[lang].get(
            items.get("type", "string"), items.get("type", "string")
        )
        if lang == "rust":
            return f"Vec<{base_type}>"
        return f"{TYPE_MAPPINGS[lang]['array']}{base_type}"

    if field_type == "map_string_string_array":
        return TYPE_MAPPINGS[lang]["map_string_string_array"]

    return TYPE_MAPPINGS[lang].get(field_type, field_type)


def get_json_tag(field: Dict, lang: str) -> str:
    """Get JSON tag for the field."""
    if lang == "go":
        json_key = field.get("json", field["name"].lower())
        omitempty = ",omitempty" if field.get("optional") else ""
        return f'`json:"{json_key}{omitempty}"`'
    return ""


def generate_go_type(task_name: str, task_def: Dict, definitions: Dict) -> str:
    """Generate Go type definitions."""
    output_type = task_def["outputType"]
    source_type = task_def.get("sourceType", "json_string")
    lines = [f"type {output_type} struct {{"]

    for field in task_def.get("fields", []):
        field_type = get_field_type(field, definitions, "go")
        # For direct source type, use struct field name for JSON tag
        # For json_string source type, use the sourceField for JSON tag
        if source_type == "direct":
            json_key = field.get("json", field["name"])
        else:
            json_key = field.get("json", field["name"].lower())
        omitempty = ",omitempty" if field.get("optional") else ""
        lines.append(f'    {field["name"]} {field_type} `json:"{json_key}{omitempty}"`')

    lines.append("}")
    lines.append("")

    # Generate definitions
    for def_name, def_def in task_def.get("definitions", {}).items():
        lines.append(f"type {def_name} struct {{")
        for field in def_def.get("fields", []):
            field_type = get_field_type(field, definitions, "go")
            json_key = field.get("json", field["name"].lower())
            omitempty = ",omitempty" if field.get("optional") else ""
            lines.append(
                f'    {field["name"]} {field_type} `json:"{json_key}{omitempty}"`'
            )
        lines.append("}")
        lines.append("")

    return "\n".join(lines)


def generate_python_type(task_name: str, task_def: Dict, definitions: Dict) -> str:
    """Generate Python type definitions using dataclasses."""
    output_type = task_def["outputType"]
    lines = [
        "from dataclasses import dataclass",
        "from typing import List, Dict, Optional",
        "",
        "",
        f"@dataclass",
        f"class {output_type}:",
    ]

    for field in task_def.get("fields", []):
        field_type = get_field_type(field, definitions, "python")
        optional = "Optional[" if field.get("optional") else ""
        default = " = None" if field.get("optional") else ""
        lines.append(f"    {field['name'].lower()}: {optional}{field_type}{default}")

    lines.append("")

    # Generate definitions
    for def_name, def_def in task_def.get("definitions", {}).items():
        lines.append(f"@dataclass")
        lines.append(f"class {def_name}:")
        for field in def_def.get("fields", []):
            field_type = get_field_type(field, definitions, "python")
            optional = "Optional[" if field.get("optional") else ""
            default = " = None" if field.get("optional") else ""
            lines.append(
                f"    {field['name'].lower()}: {optional}{field_type}{default}"
            )
        lines.append("")

    return "\n".join(lines)


def generate_ts_type(task_name: str, task_def: Dict, definitions: Dict) -> str:
    """Generate TypeScript type definitions."""
    output_type = task_def["outputType"]
    lines = [f"export interface {output_type} {{"]

    for field in task_def.get("fields", []):
        field_type = get_field_type(field, definitions, "ts")
        optional = "?" if field.get("optional") else ""
        lines.append(f"    {field['name']}{optional}: {field_type};")

    lines.append("}")
    lines.append("")

    # Generate definitions
    for def_name, def_def in task_def.get("definitions", {}).items():
        lines.append(f"export interface {def_name} {{")
        for field in def_def.get("fields", []):
            field_type = get_field_type(field, definitions, "ts")
            optional = "?" if field.get("optional") else ""
            lines.append(f"    {field['name'].lower()}{optional}: {field_type};")
        lines.append("}")
        lines.append("")

    return "\n".join(lines)


def generate_rust_type(task_name: str, task_def: Dict, definitions: Dict) -> str:
    """Generate Rust type definitions."""
    output_type = task_def["outputType"]
    lines = [f"#[derive(Deserialize, Serialize)]", f"pub struct {output_type} {{"]

    for field in task_def.get("fields", []):
        field_type = get_field_type(field, definitions, "rust")
        json_key = field.get("json", field["name"])
        optional = (
            ', skip_serializing_if = "Option::is_none"' if field.get("optional") else ""
        )
        lines.append(f'    #[serde(rename = "{json_key}"{optional})]')
        lines.append(f"    pub {field['name'].lower()}: {field_type},")

    lines.append("}")
    lines.append("")

    # Generate definitions
    for def_name, def_def in task_def.get("definitions", {}).items():
        lines.append(f"#[derive(Deserialize, Serialize)]")
        lines.append(f"pub struct {def_name} {{")
        for field in def_def.get("fields", []):
            field_type = get_field_type(field, definitions, "rust")
            json_key = field.get("json", field["name"].lower())
            optional = (
                ', skip_serializing_if = "Option::is_none"'
                if field.get("optional")
                else ""
            )
            lines.append(f'    #[serde(rename = "{json_key}"{optional})]')
            lines.append(f"    pub {field['name'].lower()}: {field_type},")
        lines.append("}")
        lines.append("")

    return "\n".join(lines)


def generate_go_converter(task_def: Dict) -> str:
    """Generate Go AsXxx() method for TaskResponse."""
    output_type = task_def["outputType"]
    source_field = task_def.get("sourceField", "")
    source_type = task_def.get("sourceType", "json_string")

    if source_type == "direct":
        field0 = task_def["fields"][0]
        field1 = task_def["fields"][1]
        # Use ADP field names to read from ExecutionMetaData
        adp_field0 = field0.get("sourceField", field0["name"].lower())
        adp_field1 = field1.get("sourceField", field1["name"].lower())

        code = (
            "func (r *TaskResponse) As"
            + output_type
            + "() (*"
            + output_type
            + ", error) {\n"
        )
        code += "    if r.ExecutionMetaData == nil {\n"
        code += '        return nil, errors.New("no execution metadata")\n'
        code += "    }\n\n"
        code += "    output := &" + output_type + "{\n"
        code += (
            "        "
            + field0["name"]
            + ': r.ExecutionMetaData["'
            + adp_field0
            + '"].(string),\n'
        )
        code += (
            "        "
            + field1["name"]
            + ': r.ExecutionMetaData["'
            + adp_field1
            + '"].(string),\n'
        )
        code += "    }\n"
        code += "    return output, nil\n"
        code += "}\n"
        return code

    # Default: parse JSON string from metadata
    code = (
        "func (r *TaskResponse) As"
        + output_type
        + "() (*"
        + output_type
        + ", error) {\n"
    )
    code += "    if r.ExecutionMetaData == nil {\n"
    code += '        return nil, errors.New("no execution metadata")\n'
    code += "    }\n\n"
    code += (
        "    "
        + source_field
        + ', ok := r.ExecutionMetaData["'
        + source_field
        + '"].(string)\n'
    )
    code += "    if !ok {\n"
    code += '        return nil, errors.New("no ' + source_field + ' found")\n'
    code += "    }\n\n"
    code += "    var output " + output_type + "\n"
    code += (
        "    if err := json.Unmarshal([]byte("
        + source_field
        + "), &output); err != nil {\n"
    )
    code += (
        '        return nil, fmt.Errorf("failed to parse '
        + output_type
        + ': %w", err)\n'
    )
    code += "    }\n\n"
    code += "    return &output, nil\n"
    code += "}\n"
    return code


def main():
    try:
        parser = argparse.ArgumentParser(description="Generate ADP output types")
        parser.add_argument(
            "--lang",
            choices=["go", "python", "ts", "rust"],
            required=True,
            help="Target language",
        )
        parser.add_argument("--output", required=True, help="Output directory")
        parser.add_argument(
            "--schema", default="specs/tasks/outputs.json", help="Path to outputs.json"
        )
        parser.add_argument(
            "--package",
            default="adp",
            help="Package name for Go/TS, module name for Python",
        )
        args = parser.parse_args()

        print(
            f"Args: lang={args.lang}, output={args.output}, schema={args.schema}",
            file=sys.stderr,
        )

        # Load schema
        with open(args.schema, "r") as f:
            schema = json.load(f)

        definitions = schema.get("definitions", {})

        # Generate types
        all_types: List[str] = []
        converters: List[str] = []

        for task_name, task_def in schema["tasks"].items():
            if args.lang == "go":
                all_types.append(generate_go_type(task_name, task_def, definitions))
                converters.append(generate_go_converter(task_def))
            elif args.lang == "python":
                all_types.append(generate_python_type(task_name, task_def, definitions))
            elif args.lang == "ts":
                all_types.append(generate_ts_type(task_name, task_def, definitions))
            elif args.lang == "rust":
                all_types.append(generate_rust_type(task_name, task_def, definitions))

        # Write output
        output_dir = Path(args.output)
        output_dir.mkdir(parents=True, exist_ok=True)

        content = ""
        if args.lang == "go":
            output_file = output_dir / "output_types.go"
            content = """package adp

import (
    "encoding/json"
    "errors"
    "fmt"
)

"""
            content += "\n\n".join(all_types)
            content += "\n\n// TaskResponse converters\n\n"
            content += "\n\n".join(converters)

        elif args.lang == "python":
            output_file = output_dir / "output_types.py"
            content = '''"""ADP Output Types - Auto-generated from outputs.json"""

'''
            content += "\n\n".join(all_types)
            content += "\n\n# TaskResponse converters need to be added manually\n"

        elif args.lang == "ts":
            output_file = output_dir / "output_types.ts"
            content = "// Auto-generated from outputs.json\n\n"
            content += "\n\n".join(all_types)

        elif args.lang == "rust":
            output_file = output_dir / "output_types.rs"
            content = "// Auto-generated from outputs.json\n\n"
            content += "use serde::{Deserialize, Serialize};\nuse std::collections::HashMap;\n\n"
            content += "\n\n".join(all_types)

        output_file.write_text(content)
        print(f"Generated {args.lang} types: {output_file}")
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        import traceback

        traceback.print_exc()
        sys.exit(1)

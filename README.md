# proto-go-course
Protobuf w/ Golang - course from Udemy

## Advantages

- Typed
- Generate code (Python, JS, C++, Go, etc.)
- Focus on optimizing types
- Schema evolution
- Comments (documentation)
- Data serialized in binary, not strings

## Disadvantages

- Serialized data is not readable
- Less support than JSON and XML

## Basics

```go
syntax = "proto3"
message Account {
    uint32 id = 1;
    string name = 2;
    bool is_verified = 3;
}
```

- When a field doesn't receive a value, it´s not serialized
- Unset fields will be populated by default values
  - int32, int64, sint32, sint64
  - uint32, uint64
  - fixed32, fixed64, sfixed32, sfixed64 (fixed amount of bytes)
  - float, double
    - Default: 0
  - bool (true or false)
    - Default: false
  - string (UTF8 or ASCII encoded )
    - Default: ""
  - bytes (arbitrary length byte sequence)
    - Default: empty bytes

### Tags are most important, used for serialization (not names)

- Smallest tag: 1
- Smaller the tag is, smaller the payload
  - Use smaller tags for most populates fields in the schema (1-15)

### Repeated fields

- Keyword: `repeated <type> <name> = <tag>;`
- Value: any number of elements
- Default: empty list

### Enums

- Keyword: `enum`
- Default value: first value
- **First tag should be 0**

### Comments

- // comment
- /\* comment \*/

## Running Address Book (self-guided exercise)

Run `make practice` so the binaries are generated in `bin/`.

To add an entry in the address book run:

```bash
./bin/add_person addressbook.data
```

To see all entries run:

```bash
./bin/list_people addressbook.data
```

## Data Evolution

- Backwards compatible
  - Don't change field tags
  - It's ok to add new fields (ignored by older versions)
  - Use reserved tags (when removing fields - makes field tag unusable)
  - Check for type compatibility

### Renaming fields

Field naming matters for the code, not for serialization/deserialization. The field tag matters.

```go
// Old
syntax = "proto3";

message Account {
    uint32 id = 1;
    string first_name = 2;
}
```

```go
// New
syntax = "proto3";

message Account {
    uint32 id = 1;
    string alias = 2;
}
```

### Removing fields

Use the `reserved` so the tag can't be used again. This also can be done for the field name.

```go
// Old
syntax = "proto3";

message Account {
    uint32 id = 1;
    string first_name = 2;
}
```

```go
// New
syntax = "proto3";

message Account {
    reserved 2;
    reserved "first_name"; //optional
    uint32 id = 1;
}
```

#### Different uses of `reserved` keywords

- Can't reserve field names and field types in a single line/keyword
- It's possible to use ranges of tags

```go
syntax = "proto3";

message Account {
    reserved 2, 15, 9 to 11;
    reserved "first_name", "last_name";
    uint32 id = 1;
}
```

### Default values

**Positive:**

- Enables forwards and backwards compatibility
- Prevents `null` values

**Negative:**

- Cannot differentiate missing from unset values
  - Not give meaning for default values
  - `if` or `switch` to reject default values

## Advanced

### Decode Raw

Used for decoding binaries into tag-value text. Reads from standard input and writes on standart output.
Running example:

`cat advanced/protoc/simple.bin | protoc --decode_raw`

### Decode

Used for decoding binaries into keyword-value text. Reads from standard input and writes on standart output.
Running example:

`cat advanced/protoc/simple.bin | protoc --decode=Simple advanced/protoc/simple.proto`

When there is a package defined in the `.proto` file add the package name: `... --decode=<pkg_name>.Simple ...`
Example:

`cat advanced/protoc/simple.bin | protoc --decode=simple.Simple advanced/protoc/simple.proto`

### Encode

It's the opposite of Decode: receives a keyword-value text input and translates it into binary.
Running example:

`cat advanced/protoc/simple.txt | protoc --encode=Simple advanced/protoc/simple.proto > advanced/protoc/simple.pb`

Just like Decode, if there's a package defined it should be added in the command: `... --encode=<pkg_name>.Simple ...`


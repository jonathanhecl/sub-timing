# 🕐 Sub-Timing
A cross-platform SRT/SSA subtitle timing correction tool.

## Features

- Support SRT/SSA subtitle files
- Move subtitle timing based on first line appearance time
- Shift subtitle timing by a specified duration
- Adjust the subtitle timing by a specified duration (between the first line and the last line)

## Usage

```bash
SubTiming [mode] -s=source.srt [-o=output.srt] [-f="00:00:00.000"] [-d="+00:00:00.000"] [-l="00:00:00.000"]
```

## Parameters

- `[mode]`: Mode to perform:
  - `move`: Move subtitle timing based on first line appearance time
  - `shift`: Shift subtitle timing by a specified duration
  - `adjust`: Adjust the subtitle timing by a specified duration (between the first line and the last line).
- `-s`: Source subtitle file
- `-o`: Output subtitle file (default: source_modified.srt)
- `-d`: Duration to shift by (for shift mode, can be positive or negative) 
- `-f`: First line appearance time (for move and adjust mode)
- `-l`: Last line appearance time (for adjust mode)

## Examples

```bash
# Move subtitle timing based on first line appearance time
SubTiming move -s=source.srt -f="00:00:00.000"

# Shift subtitle timing by a specified duration
SubTiming shift -s=source.srt -d="-00:00:00.000"

# Adjust the subtitle timing by a specified duration (between the first line and the last line)
SubTiming adjust -s=source.srt -f="00:00:00.000" -l="00:00:00.000"
```

## License

[MIT](LICENSE)

## Release

[Release](https://github.com/jonathanhecl/sub-timing/releases)
# 🕐 Sub-Timing
A cross-platform SRT/SSA subtitle timing correction tool.

## Usage

```bash
sub-timing [mode] -s=source.srt [-o=output.srt] [-f=00:00:00.000] [-d=00:00:00.000]
```

## Parameters

- `[mode]`: Mode to perform:
  - `move`: Move subtitle timing based on first line appearance time
  - `shift`: Shift subtitle timing by a specified duration
- `-s`: Source subtitle file
- `-o`: Output subtitle file (default: source_modified.srt)
- `-f`: First line appearance time (for move mode)
- `-d`: Duration to shift by (for shift mode)


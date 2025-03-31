# 🕐 Sub-Timing
A cross-platform SRT/SSA subtitle timing correction tool.

## Usage

```bash
sub-timing [mode] -s=source.srt [-o=output.srt] [-f=00:00:00.000] [-d=00:00:00.000] [-l=00:00:00.000]
```

## Parameters

- `[mode]`: Mode to perform:
  - `move`: Move subtitle timing based on first line appearance time
  - `shift`: Shift subtitle timing by a specified duration
  - `adjust`: Adjust subtitle timing by a specified duration (for both start and end times)
- `-s`: Source subtitle file
- `-o`: Output subtitle file (default: source_modified.srt)
- `-d`: Duration to shift by (for shift mode)
- `-f`: First line appearance time (for move and adjust mode)
- `-l`: Last line appearance time (for adjust mode)


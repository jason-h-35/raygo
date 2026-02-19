# Raygo: A Golang CPU-Based 3D Ray Tracer

## Requirements
- Go `1.21.x` (recommended for local parity with CI)

## Quick Start
```bash
make check
make run
```

The demo writes:
- `out.ppm`
- `out.png`
- `out.jpg`

## Notes
- Internal color math uses linear `float32` RGB.
- Image export currently clamps to display range `[0,1]` and quantizes at output boundaries.

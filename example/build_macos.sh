#!/bin/bash

# Build script for macOS systray example
echo "Building systray example for macOS..."

# Set required environment variables
export CGO_ENABLED=1

# Build the binary
echo "Step 1: Building binary..."
go build -o systray_example main.go

if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

echo "Step 2: Creating app bundle..."
# Create app bundle structure
mkdir -p SystrayExample.app/Contents/{MacOS,Resources}

# Copy the binary
cp systray_example SystrayExample.app/Contents/MacOS/

# Create Info.plist
cat > SystrayExample.app/Contents/Info.plist << EOF
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>systray_example</string>
    <key>CFBundleIdentifier</key>
    <string>com.example.systray</string>
    <key>CFBundleName</key>
    <string>SystrayExample</string>
    <key>CFBundleVersion</key>
    <string>1.0</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>NSHighResolutionCapable</key>
    <string>True</string>
    <key>LSUIElement</key>
    <string>1</string>
</dict>
</plist>
EOF

echo "Build complete!"
echo ""
echo "To run the example:"
echo "1. Method 1 (App Bundle): open SystrayExample.app"
echo "2. Method 2 (Direct):     ./systray_example"
echo "3. Method 3 (Go Run):     CGO_ENABLED=1 go run main.go"
echo ""
echo "Look for the system tray icon in the top menu bar (near the clock)."
echo "If you don't see it, try running from Terminal and check for errors."
# Release Checklist - v0.0.1

## ✅ Completed

- [x] Create repository github.com/moblang/mob
- [x] Implement compiler pipeline
- [x] Create CLI with all commands
- [x] Write comprehensive documentation
- [x] Create test suite
- [x] Add CI/CD workflows
- [x] Create installation script
- [x] Add CHANGELOG.md
- [x] Push all code to GitHub
- [x] Create tag v0.0.1
- [x] Push tag to trigger release

## 🔄 In Progress

- [ ] GitHub Actions building binaries
- [ ] Creating GitHub Release
- [ ] Uploading binaries to release

## ⏳ After Release Completes

Check the following:

### 1. Verify Release Page
Visit: https://github.com/moblang/mob/releases/tag/v0.0.1

Expected:
- [ ] Release is marked as "Latest release"
- [ ] Binary mob-linux-amd64 is attached
- [ ] Binary mob-linux-arm64 is attached
- [ ] Release notes are displayed
- [ ] Not marked as pre-release or draft

### 2. Test Installation

```bash
# Test the install script
curl -fsSL https://raw.githubusercontent.com/moblang/mob/main/install.sh | bash

# Verify installation
mob version

# Test Hello World
echo 'print("Hello World!")' > test.mob
mob run test.mob
```

### 3. Test Direct Download

```bash
# Download AMD64
wget https://github.com/moblang/mob/releases/download/v0.0.1/mob-linux-amd64
chmod +x mob-linux-amd64
./mob-linux-amd64 version

# Download ARM64
wget https://github.com/moblang/mob/releases/download/v0.0.1/mob-linux-arm64
chmod +x mob-linux-arm64
./mob-linux-arm64 version
```

### 4. Verify GitHub Actions

Visit: https://github.com/moblang/mob/actions

Expected:
- [ ] "Release" workflow completed successfully
- [ ] "CI" workflow passed for the tag
- [ ] No red ✖️ marks

### 5. Test Platform Support

**On x86_64 (amd64):**
```bash
uname -m  # Should output: x86_64
mob version
```

**On ARM64:**
```bash
uname -m  # Should output: aarch64
mob version
```

## 📢 Announcement Checklist

After release is verified, prepare announcement:

- [ ] Write blog post about moblang
- [ ] Create tweet/social post
- [ ] Post on Reddit (r/programming, r/golang)
- [ ] Share on Hacker News
- [ ] Update personal website/blog
- [ ] Add to Awesome Compilers list

## 🎯 Next Release (v0.1.0)

Planned features:
- [ ] Variable declarations
- [ ] Expression support
- [ ] Type system
- [ ] Control flow
- [ ] User-defined functions
- [ ] Class implementation

## 🐛 Troubleshooting

### Release not appearing
1. Check GitHub Actions status
2. Check for authentication issues
3. Verify tag was pushed: `git tag -l`

### Binaries not downloading
1. Check if release is published (not draft)
2. Verify file permissions
3. Check download URLs in release page

### Installation script fails
1. Check if curl/wget is installed
2. Verify write permissions in ~/.local/bin
3. Check network connectivity

## 📞 Support

If you encounter issues:
- Check Actions: https://github.com/moblang/mob/actions
- Check Issues: https://github.com/moblang/mob/issues
- Check Discussions: https://github.com/moblang/mob/discussions

import re

FILE = 'backend/internal/search/dataset.go'

ICONS = {
    'figma': 'https://upload.wikimedia.org/wikipedia/commons/3/33/Figma-logo.svg',
    'canva': 'https://upload.wikimedia.org/wikipedia/commons/0/08/Canva_icon_2021.svg',
    'penpot': 'https://upload.wikimedia.org/wikipedia/commons/3/3f/Penpot_logo.svg',
    'vscode': 'https://upload.wikimedia.org/wikipedia/commons/9/9a/Visual_Studio_Code_1.35_icon.svg',
    'intellij': 'https://upload.wikimedia.org/wikipedia/commons/9/9c/IntelliJ_IDEA_Icon.svg',
    'sublime': 'https://upload.wikimedia.org/wikipedia/commons/7/79/Sublime_Text_Logo.svg',
    'notion': 'https://upload.wikimedia.org/wikipedia/commons/e/e9/Notion-logo.svg',
    'obsidian': 'https://upload.wikimedia.org/wikipedia/commons/1/10/2023_Obsidian_logo.svg',
    'chrome': 'https://upload.wikimedia.org/wikipedia/commons/8/87/Google_Chrome_icon_%282011%29.png',
    'firefox': 'https://upload.wikimedia.org/wikipedia/commons/a/a0/Firefox_logo%2C_2019.svg',
    'safari': 'https://upload.wikimedia.org/wikipedia/commons/5/52/Safari_browser_logo.svg',
    'arc': 'https://upload.wikimedia.org/wikipedia/commons/8/87/Arc_logo.svg',
    'blender': 'https://upload.wikimedia.org/wikipedia/commons/0/0c/Blender_logo_no_text.svg',
    'maya': 'https://upload.wikimedia.org/wikipedia/commons/f/fb/Autodesk_Maya_2013_logo.png',
    'slack': 'https://upload.wikimedia.org/wikipedia/commons/d/d5/Slack_icon_2019.svg',
    'discord': 'https://upload.wikimedia.org/wikipedia/commons/9/90/Discord_logo_2021.svg',
    'teams': 'https://upload.wikimedia.org/wikipedia/commons/c/c9/Microsoft_Office_Teams_%282018%E2%80%93present%29.svg',
    'zoom': 'https://upload.wikimedia.org/wikipedia/commons/7/7b/Zoom_Communications_Logo.svg',
    'spotify': 'https://upload.wikimedia.org/wikipedia/commons/2/26/Spotify_logo_with_text.svg',
}

with open(FILE, 'r', encoding='utf-8') as f:
    content = f.read()

for key, url in ICONS.items():
    pattern = r'Icon:\s*"' + re.escape(key) + r'"'
    replacement = f'Icon:         "{url}"'
    content = re.sub(pattern, replacement, content)

with open(FILE, 'w', encoding='utf-8') as f:
    f.write(content)
print("Updated icons in dataset.go")

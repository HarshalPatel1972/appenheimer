export const KNOWN_ICONS: Record<string, string> = {
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
	'illustrator': 'https://upload.wikimedia.org/wikipedia/commons/f/fb/Adobe_Illustrator_CC_icon.svg',
	'photoshop': 'https://upload.wikimedia.org/wikipedia/commons/a/af/Adobe_Photoshop_CC_icon.svg',
	'premiere': 'https://upload.wikimedia.org/wikipedia/commons/4/40/Adobe_Premiere_Pro_CC_icon.svg',
	'ableton': 'https://upload.wikimedia.org/wikipedia/commons/1/1a/Ableton_Live_11_icon.png',
};

export function getIconUrl(icon: string | undefined, name: string): string {
	if (!icon) return '';
	
	// If it's already an absolute URL (from a fresh backend), use it
	if (icon.startsWith('http')) return icon;
	
	// Check our hardcoded fallback dictionary for stale backends
	const cleanSlug = icon.toLowerCase().replace(/[^a-z0-9]/g, '');
	if (KNOWN_ICONS[cleanSlug]) return KNOWN_ICONS[cleanSlug];
	
	// If the backend returns a short slug and we don't have it hardcoded, use Google Favicons as the final, unblockable fallback
	return `https://www.google.com/s2/favicons?domain=${cleanSlug}.com&sz=128`;
}

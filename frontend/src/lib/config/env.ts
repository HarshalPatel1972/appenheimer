const apiUrl = import.meta.env.VITE_API_URL;

if (!apiUrl) {
	console.error('CRITICAL: VITE_API_URL is missing! Requests will fail or hit localhost/vercel directly. Please add it to Vercel Environment Variables and redeploy.');
}

export const ENV = {
	API_URL: apiUrl || 'https://appenheimer.onrender.com',
};

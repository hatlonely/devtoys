import { purgeCss } from 'vite-plugin-tailwind-purgecss';
import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';


export default defineConfig({
	server: {
	fs: {
		allow: [
			'wailsjs'
		],
	},
	},
	plugins: [
		sveltekit(), 
		purgeCss(),
	]
});

import { loadSetting } from "$lib/store/store"

export const prerender = true
export const ssr = false

export function load() {
    loadSetting()
    
    return {};
}

import { writable } from 'svelte/store';
import { GetSetting } from '$lib/wailsjs/go/devtoys/App';

const toolItems = [
    {
        icon: 'base64',
        title: 'Base64 文本编解码',
        description: 'Base64 文本编解码工具',
        href: '/base64-text'
    },
    {
        icon: 'timestamp',
        title: 'Unix 时间戳',
        description: 'Unix 时间戳工具',
        href: '/unix-timestamp'
    },
    {
        icon: 'escape',
        title: '字符串转义',
        description: '字符串、HTML、URL、URL Query 转义工具',
        href: '/escape'
    },
    {
        icon: 'text',
        title: '文本二进制转换',
        description: '查看文本的二进制、八进制、十六进制标识',
        href: '/text-binary-conversion'
    },
    {
        icon: 'number',
        title: '数值进制转换',
        description: '数值二进制、八进制、十进制、十六进制转换',
        href: '/number-base-conversion'
    },
    {
        icon: 'hash',
        title: '哈希散列',
        description: 'MD5、SHA1、SHA256、SHA512、HMAC 哈希散列工具',
        href: '/hash'
    },
    {
        icon: 'string',
        title: '字符串转换',
        description: '字符串大小写转换、驼峰转换、下划线转换、中划线转换',
        href: '/string-conversion'
    },
    {
        icon: 'https',
        title: 'https 证书',
        description: 'https 证书解析工具',
        href: '/ssl-certificate'
    },
    {
        icon: 'password',
        title: '密码生成',
        description: '密码生成工具',
        href: '/password-generator'
    },
    {
        icon: 'json',
        title: 'JSON 格式化',
        description: 'JSON 格式化工具',
        href: '/json-formatter',
    },
    {
        icon: 'sql',
        title: 'SQL 格式化',
        description: 'SQL 格式化工具',
        href: '/sql-formatter',
    },
    {
        icon: 'yaml',
        title: 'JSON YAML 转换',
        description: 'JSON YAML 转换工具',
        href: '/json-to-yaml',
    },
    {
        icon: 'http',
        title: 'HTTP 客户端',
        description: 'HTTP 转换工具',
        href: '/http-client',
    }
];

export const store = {
    toolItems,
};

export const activeToolItemHref = writable('/');
export const title = writable('所有工具');
export const setting = writable(Object());

export async function loadSetting() {
    const appSetting = await GetSetting();
    setting.set(appSetting);
}

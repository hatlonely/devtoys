import { writable } from 'svelte/store';

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
    }
];

export const store = {
    toolItems,
};

export const activeToolItemHref = writable('');

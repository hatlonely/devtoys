<script context="module" lang="ts">
	import hljs from 'highlight.js/lib/core';
	import javascript from 'highlight.js/lib/languages/javascript';
	import json from 'highlight.js/lib/languages/json';
	import cpp from 'highlight.js/lib/languages/cpp';
	import python from 'highlight.js/lib/languages/python';
	import java from 'highlight.js/lib/languages/java';
	import bash from 'highlight.js/lib/languages/bash';
	import css from 'highlight.js/lib/languages/css';
	import markdown from 'highlight.js/lib/languages/markdown';
	import nginx from 'highlight.js/lib/languages/nginx';
	import php from 'highlight.js/lib/languages/php';
	import sql from 'highlight.js/lib/languages/sql';
	import xml from 'highlight.js/lib/languages/xml';
	import yaml from 'highlight.js/lib/languages/yaml';
	import typescript from 'highlight.js/lib/languages/typescript';
	import go from 'highlight.js/lib/languages/go';
	import kotlin from 'highlight.js/lib/languages/kotlin';
	import rust from 'highlight.js/lib/languages/rust';
	import swift from 'highlight.js/lib/languages/swift';
	import scala from 'highlight.js/lib/languages/scala';
	import ruby from 'highlight.js/lib/languages/ruby';
	import perl from 'highlight.js/lib/languages/perl';
	import lua from 'highlight.js/lib/languages/lua';
	import 'highlight.js/styles/monokai.css';

	hljs.registerLanguage('javascript', javascript);
	hljs.registerLanguage('json', json);
	hljs.registerLanguage('cpp', cpp);
	hljs.registerLanguage('python', python);
	hljs.registerLanguage('java', java);
	hljs.registerLanguage('bash', bash);
	hljs.registerLanguage('css', css);
	hljs.registerLanguage('markdown', markdown);
	hljs.registerLanguage('nginx', nginx);
	hljs.registerLanguage('php', php);
	hljs.registerLanguage('sql', sql);
	hljs.registerLanguage('xml', xml);
	hljs.registerLanguage('yaml', yaml);
	hljs.registerLanguage('typescript', typescript);
	hljs.registerLanguage('go', go);
	hljs.registerLanguage('kotlin', kotlin);
	hljs.registerLanguage('rust', rust);
	hljs.registerLanguage('swift', swift);
	hljs.registerLanguage('scala', scala);
	hljs.registerLanguage('ruby', ruby);
	hljs.registerLanguage('perl', perl);
	hljs.registerLanguage('lua', lua);

	const highlight = (code: any, syntax: any) => {
		if (syntax) {
			return hljs.highlight(code, {
				language: syntax
			}).value;
		} else {
			return hljs.highlightAuto(code).value;
		}
	};
</script>

<script lang="ts">
	import { CodeJar } from '@novacbn/svelte-codejar';

	export let value = '';
	export let syntax = '';
</script>

<CodeJar class="hljs" {syntax} {highlight} {value} />

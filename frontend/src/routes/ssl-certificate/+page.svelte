<script lang="ts">
	import { SSLCertificate } from '$lib/wailsjs/go/devtoys/App';
	import { Title, Textarea, TextViewer, DataTable } from '$lib';
	import '@fontsource/roboto-mono';

	let text = '';
	let warning: any = '';
	let result = '';
	let subjectTable: any = {};
	let issuerTable: any = {};
	let certificateTable: any = {};

	async function calculate() {
		warning = '';
		if (!text) {
			result = '';
			return;
		}
		try {
			let res = await SSLCertificate({
				Text: text
			});
			result = res;
			subjectTable = calculateSubjectTable(res);
			issuerTable = calculateIssuerTable(res);
			certificateTable = calculateCertificateTable(res);
		} catch (err) {
			warning = err;
		}
	}

	$: text, calculate();

	function calculateSubjectTable(res: any) {
		return {
			title: '域名信息',
			names: ['label', 'value'],
			rows: [
				{
					label: '国家',
					value: res.Subject.Country ? res.Subject.Country.join(', ') : ''
				},
				{
					label: '组织',
					value: res.Subject.Organization ? res.Subject.Organization.join(', ') : ''
				},
				{
					label: '组织单位',
					value: res.Subject.OrganizationUnit ? res.Subject.OrganizationUnit.join(', ') : ''
				},
				{
					label: '域名',
					value: res.Subject.CommonName
				},
				{
					label: '省份',
					value: res.Subject.Province ? res.Subject.Province.join(', ') : ''
				},
				{
					label: '地区',
					value: res.Subject.Locality ? res.Subject.Locality.join(', ') : ''
				},
				{
					label: '街道地址',
					value: res.Subject.StreetAddress ? res.Subject.StreetAddress.join(', ') : ''
				},
				{
					label: '邮政编码',
					value: res.Subject.PostalCode ? res.Subject.PostalCode.join(', ') : ''
				}
			]
		};
	}

	function calculateIssuerTable(res: any) {
		return {
			title: '颁发者信息',
			names: ['label', 'value'],
			rows: [
				{
					label: '国家',
					value: res.Issuer.Country ? res.Issuer.Country.join(', ') : ''
				},
				{
					label: '组织',
					value: res.Issuer.Organization ? res.Issuer.Organization.join(', ') : ''
				},
				{
					label: '组织单位',
					value: res.Issuer.OrganizationUnit ? res.Issuer.OrganizationUnit.join(', ') : ''
				},
				{
					label: '颁发机构',
					value: res.Issuer.CommonName
				}
			]
		};
	}

	function calculateCertificateTable(res: any) {
		return {
			title: '证书信息',
			names: ['label', 'value'],
			rows: [
				{ label: '证书序列号', value: res.Certificate.SerialNumber },
				{ label: '证书生效时间', value: res.Certificate.NotBefore },
				{ label: '证书失效时间', value: res.Certificate.NotAfter },
				{ label: '签名算法', value: res.Certificate.SignatureAlgorithm },
				{ label: '公钥算法', value: res.Certificate.PublicKeyAlgorithm },
				{
					label: '公钥',
					value: res.Certificate.PublicKey,
					_meta: {
						value: {
							code: true
						}
					}
				},
				{ label: '颁发者公钥标识', value: res.Certificate.AuthorityKeyId },
				{ label: '使用者公钥标识', value: res.Certificate.SubjectKeyId },
				{ label: 'MD5', value: res.Certificate.Md5 },
				{ label: 'SHA1', value: res.Certificate.Sha1 },
				{ label: 'SHA256', value: res.Certificate.Sha256 },
				{
					label: 'DNS 域名',
					value: res.Certificate.DNSNames ? res.Certificate.DNSNames.join(', ') : ''
				},
				{
					label: 'CRL 分发点',
					value: res.Certificate.CRLDistributionPoints
						? res.Certificate.CRLDistributionPoints.join(', ')
						: ''
				},
				{
					label: '颁发者证书 URL',
					value: res.Certificate.IssuingCertificateURL
						? res.Certificate.IssuingCertificateURL.join(', ')
						: ''
				},
				{
					label: 'OCSP 服务器地址',
					value: res.Certificate.OCSPServer ? res.Certificate.OCSPServer.join(', ') : ''
				}
			]
		};
	}
</script>

<div class="w-full text-token px-6 py-3 space-y-4">
	<Title title="字符串转换" />
	<!-- <div class="w-full text-token card p-4" /> -->

	<div class="w-full text-token card p-4">
		<Textarea
			bind:value={text}
			on:input={calculate}
			title="证书"
			placeholder="输入 https 证书"
			code={true}
			{warning}
		/>
	</div>

	{#if result}
		<DataTable title={subjectTable.title} names={subjectTable.names} rows={subjectTable.rows} />
		<DataTable title={issuerTable.title} names={issuerTable.names} rows={issuerTable.rows} />
		<DataTable
			title={certificateTable.title}
			names={certificateTable.names}
			rows={certificateTable.rows}
		/>
	{/if}
</div>

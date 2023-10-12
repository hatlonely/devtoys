<script lang="ts">
	import { AppShell, AppBar } from '@skeletonlabs/skeleton';
	import 'material-icons/iconfont/material-icons.css';
</script>

<AppShell>
	<svelte:fragment>
		<div class="card p-4 w-full text-token space-y-4">
			<div class="card p-4">
				<span class="badge"><span class="material-icons-outlined">sync_alt</span></span>
				<span class="">转换</span>
				<div class="card p-4 flex justify-center items-center">
					<span>Minimal</span>
				</div>
			</div>

			<label class="label">
				<textarea
					class="textarea"
					rows="4"
					placeholder="Lorem ipsum dolor sit amet consectetur adipisicing elit."
				/>
			</label>

			<label class="label">
				<textarea
					class="textarea"
					rows="4"
					placeholder="Lorem ipsum dolor sit amet consectetur adipisicing elit."
				/>
			</label>
		</div>
	</svelte:fragment>
</AppShell>

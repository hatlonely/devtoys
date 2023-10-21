<script>
	export let value = 0;

	function decrement() {
		value--;
	}

	function increment() {
		value++;
	}
</script>

<div class="custom-number-input h-10 w-36">
	<div class="flex flex-row h-10 w-full rounded-lg relative bg-transparent mt-1">
		<button
			on:click={decrement}
			class="btn bg-gray-300 hover:text-gray-700 hover:bg-gray-400 h-full w-10 rounded-r cursor-pointer"
		>
			<span class="material-symbols-outlined"> remove </span>
		</button>
		<input
			type="number"
			class="input outline-none border-transparent rounded-none text-center w-full bg-gray-300 font-semibold text-md hover:text-black focus:text-black md:text-basecursor-default flex items-center text-gray-700 outline-none"
			name="custom-input-number"
			bind:value
		/>
		<button
			on:click={increment}
			class="btn bg-gray-300 hover:text-gray-700 hover:bg-gray-400 h-full w-10 rounded-l cursor-pointer"
		>
			<span class="material-symbols-outlined"> add </span>
		</button>
	</div>
</div>

<style>
	input[type='number']::-webkit-inner-spin-button,
	input[type='number']::-webkit-outer-spin-button {
		-webkit-appearance: none;
		margin: 0;
	}

	.custom-number-input input:focus {
		outline: none !important;
	}

	.custom-number-input button:focus {
		outline: none !important;
	}
</style>

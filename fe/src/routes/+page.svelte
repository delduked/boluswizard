<script lang='ts'>
	import { fly } from 'svelte/transition';
	import { get } from 'svelte/store';

	let promise = test()
	async function test() {
		try {
			const asdf = await fetch("https://api.boluswizard.io/health")
			if (asdf.ok) {
				console.log(asdf.ok)
				return `${asdf.ok}`
			}
		} catch (error) {
			throw error
		}
	}
	function handleClick() {
		promise = test()
	}
</script>

<div
	class="hero min-h-screen"
	style="background-image: url(https://boluswizard.io/assets/loginpage/bg.gif);"
>
	<div class="hero-overlay" />
	<div
		in:fly={{y: -5, duration: 500, delay: 500}} out:fly={{y: 5, duration: 500}}
		class="hero-content text-center text-neutral-content"
	>
		<div class="max-w-md text-white">
			<h1 class="mb-5 text-5xl font-bold">Welcome to BolusWizard.io!</h1>
			<p class="mb-5">
				Don't have an expensive insulin pump? No problem! BolusWizard calculates the amount of
				insulin you need for meals and corrections! Just enter your Insulin sensitivity factor,
				Insulin duration and BG targets and you're ready to go!
			</p>
			<a class="btn btn-primary" href="/signin">Sign In</a>
			<a class="btn btn-secondary ml-4" href="/signup">Sign Up</a>
			<button on:click={handleClick} class="btn btn-neutral" >asdf</button>
		</div>
	</div>
	{#await promise}
		<p>asdf</p>
	{:then data } 
		<div class="z-50 fixed bottom-0 m-20">
			<div class="alert alert-success ">
				<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
				<span class="pr-2">{data.Status}</span>
			</div>
		</div>
	{:catch error} 
		<div class="z-50 fixed bottom-0 m-20">
			<div class="alert alert-success ">
				<svg xmlns="http://www.w3.org/2000/svg" class="stroke-current shrink-0 h-6 w-6" fill="none" viewBox="0 0 24 24"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
				<span class="pr-2">{error}</span>
			</div>
		</div>
	{/await}
</div>

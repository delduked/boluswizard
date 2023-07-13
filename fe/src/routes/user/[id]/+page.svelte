<script lang="ts">
	import { browser } from '$app/environment'; // (was '$app/env' in a pre 1.0 SvelteKit version)

	import Nav from '../../components/nav.svelte';
	import Stats from '../../components/Stats.svelte';
	import UserInput from '../../components/UserInput.svelte';
	import { onMount } from 'svelte';
	import jwt from 'jsonwebtoken';
	import Cookies from 'js-cookie';
	import { page } from '$app/stores';

	const url = $page.url;

	let chartData = [20, 100, 50, 12, 20, 130, 45];
	let Labels = ['Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday', 'Sunday'];
	let ctx;
	let chartCanvas;

	let username;
	let auth_success;

	let slug = '';

	page.subscribe((data) => {
		slug = data.url.pathname;
	});

	onMount(async () => {
		try {
			const token = Cookies.get('authToken');
			const decoded = jwt.verify(token, 'secret') as { Username: String; Uid: string };

			if (decoded.Username == slug) {
				auth_success = true;
			}
			auth_success = false;
		} catch (error) {
			console.log(error);
			auth_success = false;
		}
	});
</script>

{#if auth_success}
	<Nav />
	<div class="w-full flex justify-center items-center flex-wrap lg:m-4 sm:m-2">
		<UserInput />

		<Stats />
	</div>
	<div class="artboard phone-2">
		<canvas bind:this={chartCanvas} />
	</div>
{:else}
	<h1>nope</h1>
{/if}

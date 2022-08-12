<script context="module" lang="ts">
	export const load = async ({ params, fetch }) => {
		const member = crewByID(params.slug);
		if (!member) {
			return {
				status: 404,
				props: {
					member: null
				}
			};
		}
		const postsUrl = `/posts/blogs-by-${member.id}.json`;
		const res = await fetch(postsUrl);
		if (res.ok) {
			return {
				props: {
					member: member,
					blogs: await res.json()
				}
			};
		}
		return {
			status: res.status,
			error: new Error(`could not load ${postsUrl}`)
		};
	};
</script>

<script lang="ts">
	import { crewByID, type Member } from '$lib/crew/crew';
	import PostGrid from '$lib/posts/postGrid.svelte';
	import type { Blogs } from '$lib/posts/posts';
	import { seo } from '$lib/seo/store';

	export let member: Member;
	export let blogs: Blogs;

	seo.reset();
	$seo.title = 'Verifa Crew: ' + member.name;
	$seo.description = member.bio ? member.bio.substring(0, 100) : '';
	$seo.image.url = member.image;
</script>

{#if member.active}
	<section>
		<div class="flex flex-col lg:gap-y-8">
			<div class="w-auto flex flex-col gap-y-8 lg:flex-row sm:gap-x-4">
				<div class="max-w-sm lg:w-1/4">
					<img src={member.image} alt={member.id} class="h-full w-full object-top object-contain" />
				</div>
				<div class="w-full lg:w-3/4">
					<h2>{member.name}</h2>
					<h4>{member.position}</h4>
					<p>{member.bio}</p>
				</div>
			</div>
			<div>
				<h3>Find me online</h3>
				<div class="flex items-center gap-x-4">
					<a href={member.linkedin} class="text-v-beige hover:text-gray-500" target="_blank">
						<span class="sr-only">LinkedIn</span>
						<svg
							class="w-12 h-12 text-v-lilac hover:text-v-black"
							fill="currentColor"
							viewBox="0 0 20 20"
							aria-hidden="true"
						>
							<path
								fill-rule="evenodd"
								d="M16.338 16.338H13.67V12.16c0-.995-.017-2.277-1.387-2.277-1.39 0-1.601 1.086-1.601 2.207v4.248H8.014v-8.59h2.559v1.174h.037c.356-.675 1.227-1.387 2.526-1.387 2.703 0 3.203 1.778 3.203 4.092v4.711zM5.005 6.575a1.548 1.548 0 11-.003-3.096 1.548 1.548 0 01.003 3.096zm-1.337 9.763H6.34v-8.59H3.667v8.59zM17.668 1H2.328C1.595 1 1 1.581 1 2.298v15.403C1 18.418 1.595 19 2.328 19h15.34c.734 0 1.332-.582 1.332-1.299V2.298C19 1.581 18.402 1 17.668 1z"
								clip-rule="evenodd"
							/>
						</svg>
					</a>
					{#if member.github}
						<a href="https://github.com/{member.github}" target="_blank">
							<svg
								class="w-12 h-12 text-v-gray hover:text-v-black"
								fill="currentColor"
								xmlns="http://www.w3.org/2000/svg"
								viewBox="0 0 24 24"
							>
								<path
									d="M10.9,2.1c-4.6,0.5-8.3,4.2-8.8,8.7c-0.5,4.7,2.2,8.9,6.3,10.5C8.7,21.4,9,21.2,9,20.8v-1.6c0,0-0.4,0.1-0.9,0.1 c-1.4,0-2-1.2-2.1-1.9c-0.1-0.4-0.3-0.7-0.6-1C5.1,16.3,5,16.3,5,16.2C5,16,5.3,16,5.4,16c0.6,0,1.1,0.7,1.3,1c0.5,0.8,1.1,1,1.4,1 c0.4,0,0.7-0.1,0.9-0.2c0.1-0.7,0.4-1.4,1-1.8c-2.3-0.5-4-1.8-4-4c0-1.1,0.5-2.2,1.2-3C7.1,8.8,7,8.3,7,7.6C7,7.2,7,6.6,7.3,6 c0,0,1.4,0,2.8,1.3C10.6,7.1,11.3,7,12,7s1.4,0.1,2,0.3C15.3,6,16.8,6,16.8,6C17,6.6,17,7.2,17,7.6c0,0.8-0.1,1.2-0.2,1.4 c0.7,0.8,1.2,1.8,1.2,3c0,2.2-1.7,3.5-4,4c0.6,0.5,1,1.4,1,2.3v2.6c0,0.3,0.3,0.6,0.7,0.5c3.7-1.5,6.3-5.1,6.3-9.3 C22,6.1,16.9,1.4,10.9,2.1z"
								/></svg
							>
						</a>
					{/if}
				</div>
			</div>
		</div>
	</section>
{:else}
	<section>
		<div class="flex flex-col gap-y-8 sm:flex-row sm:gap-x-4">
			<div class="">
				<img src={member.image} alt={member.id} class="h-48 w-full" />
			</div>
			<div>
				<h2>{member.name}</h2>
			</div>
		</div>
	</section>
{/if}
<section>
	{#if blogs.blogs.length == 0}
		<h2>No posts by author</h2>
	{:else}
		<h2>Posts by author</h2>
		<PostGrid posts={blogs.blogs} />
	{/if}
</section>

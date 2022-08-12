<script lang="ts">
	import { onMount } from 'svelte';

	let attachments: string[] = [];

	function handleAttachments(event) {
		attachments = Object.values(event.target.files).map((file) => file['name']);
	}

	onMount(() => {
		// Setup listener for when files are attached
		const inputElement = document.getElementById('attachment');
		inputElement.addEventListener('change', handleAttachments, false);
	});
</script>

<div class="max-w-4xl">
	<form
		action="https://usebasin.com/f/9f1e7bc2e780"
		method="POST"
		enctype="multipart/form-data"
		id="contact-form"
		class="flex flex-col gap-y-8"
	>
		<!-- Honeypot field to avoid spamming  -->
		<input type="hidden" name="_tunaboat" />
		<div>
			<label for="name" class="block">Name*</label>
			<input type="text" name="name" id="name" autocomplete="name" required />
		</div>
		<div>
			<label for="email" class="block">Email*</label>
			<input type="email" name="email" id="email" required />
		</div>
		<div>
			<label for="message" class="block">Tell us about yourself*</label>
			<textarea id="message" name="message" rows="4" required />
		</div>
		<div class="flex flex-col gap-y-8 md:flex-row md:gap-x-4">
			<div class="flex-1">
				<label for="linkedin" class="block">LinkedIn</label>
				<input
					type="url"
					name="linkedin"
					id="linkedin"
					placeholder="https://www.linkedin.com/in/<profile>"
					pattern="https://(www.)?linkedin.com/.*"
				/>
			</div>
			<div class="flex-1">
				<label for="github" class="block">GitHub</label>
				<input
					type="url"
					name="github"
					id="github"
					placeholder="https://github.com/<profile>"
					pattern="https://(www.)?github.com/.*"
				/>
			</div>
		</div>
		<div class="flex gap-x-4 items-center">
			<div
				class="self-start py-1 px-4 border-2 border-v-lilac bg-v-lilac text-v-white  hover:bg-violet-400"
			>
				<label for="attachment" class="cursor-pointer font-medium text-lg">
					<div class="flex items-center gap-x-1">
						<svg
							xmlns="http://www.w3.org/2000/svg"
							class="h-6 w-6"
							fill="none"
							viewBox="0 0 24 24"
							stroke="currentColor"
							stroke-width="2"
						>
							<path
								stroke-linecap="round"
								stroke-linejoin="round"
								d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"
							/>
						</svg>
						<span>Attach files</span>
					</div>
				</label>
				<input type="file" id="attachment" name="attachments[]" class="hidden" multiple />
			</div>
			<div class="-my-2 flex flex-wrap gap-x-4">
				{#each attachments as attachment}
					<span class="inline-flex items-center my-2 px-3 py-0.5 bg-v-gray">
						<p class="m-0 text-v-white">{attachment}</p>
					</span>
				{/each}
			</div>
		</div>
		<div class="flex gap-x-4 items-center">
			<input
				type="checkbox"
				id="newsletter"
				name="newsletter"
				class="border-[3px] border-v-black focus:ring-0 text-v-lilac"
			/>
			<label for="newsletter" class="block mb-0">Subscribe to our monthly newsletter</label>
		</div>
		<div>
			<button class="v-button bg-v-black" type="submit">Submit</button>
		</div>
		<p>
			By submitting this form you agree to our <a href="/privacy" class="underline"
				>Privacy Policy</a
			>
		</p>
	</form>
</div>

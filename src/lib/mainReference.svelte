<script lang="ts">
	import ButtonLink from './buttonLink.svelte';
	import ClientLogos from './clients/clientLogos.svelte';
	import Column from './column.svelte';
	import Columns from './columns.svelte';
</script>

<section>
	<Columns reverse={true}>
		<Column class="self-start">
			<div class="flex flex-col gap-y-4">
				<h2>
					‘A key and crucial part of our Continuous Integration journey was having Verifa's
					involvement’.
				</h2>
				<h4>— David Hoslett, Siemens Mobility.</h4>
			</div>
		</Column>
		<Column class="self-start">
			<div class="flex flex-col gap-y-0">
				<div>
					<h4>Why work with Verifa.</h4>
					<p>
						We are a small and dedicated crew who care about the value we deliver. We look to
						tomorrow's problems and solve them for you ahead of time.
					</p>
				</div>
				<div>
					<h4>What makes us different.</h4>
					<p>
						We don't just solve problems but teach you how to solve them, and we do it as a team.
						That's our goal and what sets us apart from other consultancy houses.
					</p>
				</div>
				<div>
					<h4>Who we have helped.</h4>
					<p>
						We have helped international teams of all sizes, across many industries like Automotive,
						Medical, Finance, Mobile and Web.
					</p>
				</div>
				<div>
					<ButtonLink href="/clients">Our Clients</ButtonLink>
				</div>
			</div>
		</Column>
	</Columns>
</section>
<section>
	<ClientLogos />
</section>

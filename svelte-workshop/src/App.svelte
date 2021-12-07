<script lang="ts">
    import TopBar from './TopBar.svelte'
	import EventCards from './Card.svelte'
</script>

<TopBar/>
<div class="page">
    <EventCards/>
</div>

<style>
    :global(*) {
        padding: 0;
        margin: 0;
    }

    :global(body) {
        background-color: white;
        padding: 0;
        margin: 0;
        position: fixed;
    }

    .page {
        margin-top: 50px;
    }
</style>
<script lang="ts">
    import { v4 as uuidv4 } from 'uuid';
    import { writableArray } from "./store.js";

    import Event from './Event.svelte';
    import Add from './Add.svelte';

    $writableArray = [...$writableArray, {
        id: uuidv4(),
        title: 'Event 1',
        description: 'First example',
        start: Date.now(),
        end: Date.now(),
    }, {
        id: uuidv4(),
        title: 'Event 2',
        description: 'Second example',
        start: Date.now(),
        end: Date.now(),
    }, {
        id: uuidv4(),
        title: 'Event 3',
        description: 'Third example',
        start: Date.now(),
        end: Date.now(),
    }]

</script>

<div class="events-container">
    {#each $writableArray as event}
        <Event event={event}/>
    {/each}
    <Add/>
</div>

<style>
    .events-container {
        display: flex;
        flex-direction: row;
        flex-flow: wrap;
    }
</style>
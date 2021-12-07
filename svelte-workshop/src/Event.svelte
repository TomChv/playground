<script lang="ts" context="module">
    import type { v4 as uuidv4 }  from 'uuid';

    export type EventOptions = {
        id: typeof uuidv4;
        title: string;
        description: string;
        start: Date;
        end: Date;
    }
</script>

<script lang="ts">
    import {writableArray} from "./store.js";


    /** Month definition */
    const months: string[] = [
        "January", "February", "March", "April",
        "May", "June", "July", "August",
        "September", "October", "November", "December"
    ];

    /** Days definition */
    const days: string[] = [
        "Monday", "Tuesday", "Wednesday",
        "Thursday", "Friday", "Saturday", "Sunday"
    ];

    export let event: EventOptions;

    const start = new Date(event.start);
    const end = new Date(event.end);

    const startMonth: string = months[start.getUTCMonth()];
    const startDay: string = days[start.getDay()];
    const startHour: string = start.getHours().toString(10);
    const startMinutes: string = start.getMinutes().toString(10);


    const endMonth: string = months[end.getUTCMonth()];
    const endDay: string = days[end.getDay()];
    const endHour: string = end.getHours().toString(10);
    const endMinutes: string = end.getMinutes().toString(10);

    const fmtStartDate = [startDay, startMonth, start.getFullYear()].join(' ') + ` ${startHour}:${startMinutes}`;
    const fmtEndDate = [endDay, endMonth, end.getFullYear()].join(' ') + ` ${endHour}:${endMinutes}`;

    function deleteEvent() {
        $writableArray = $writableArray.filter((elem) => elem.id !== event.id);
    }

</script>

<div class="event-container">
    <div class="destroy-event-button-container">
        <button class="destroy-event-button-style" on:click={deleteEvent}>
            x
        </button>
    </div>

    <p class="title">
        {event.title}
    </p>
    <p class="description">
        {#if event.description.length === 0}
            No description provide
        {:else}
            {event.description}
        {/if}
    </p>
    <p>
        {fmtStartDate}
    </p>
    <p>
        {fmtEndDate}
    </p>
</div>

<style>
    .event-container {
        text-align: center;
        width: 250px;
        height: 200px;

        margin: 8px;
        padding: 5px;
        background: white;
        box-shadow: 1px 1px 3px grey;
        align-items: center;
        justify-content: center;
        position: relative;
    }

    .destroy-event-button-container {
        position: absolute;
        top: 3px;
        right: 10px;
    }

    .destroy-event-button-style {
        background: transparent;
        color: #a28585;
        font-size: 1.3em;
        transition: all .4s ease;
        border-radius: 5px;
        border: none;
        cursor: pointer;
        align-items: center;
        justify-content: center;
        margin: 0;
        padding: 0;
    }

    .destroy-event-button-style:hover {
        color: red;
    }

    .title {
        font-size: large;
        font-weight: bold;
        margin-bottom: 5px;
    }

    .description {
        margin: auto auto 15px;
        padding: 5px;
        height: 40%;
        width: 85%;
        overflow-x: hidden;
        border: #beb8b8 solid 1px;
    }

    .description::-webkit-scrollbar {
        width: 1px;
    }

    .description::-webkit-scrollbar-track {
        box-shadow: inset 0 0 6px rgba(0, 0, 0, 0.3);
    }

    .description::-webkit-scrollbar-thumb {
        outline: 1px solid #999999;
    }
</style>
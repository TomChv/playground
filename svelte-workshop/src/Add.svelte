<script lang="ts">
    import { v4 as uuidv4 } from 'uuid';
    import { writableArray } from "./store.js";

    let isOpen = false;

    function openForm() {
        isOpen = true;
    }

    function closeForm() {
        isOpen = false;
    }

    function addEvent(e: Event) {
        const {eventName, description, startDate, endDate} = e.target as any;
        const event = {
            id: uuidv4(),
            title: eventName.value,
            description: description.value,
            start: startDate.value,
            end: endDate.value,
        }
        $writableArray = [...$writableArray, event];
        closeForm();
    }
</script>


{#if isOpen}
    <div class="form-container">
        <h2 class="form-title">
            Create a new Event
        </h2>

        <form class="form-content" on:submit|preventDefault={addEvent}>
            <label class="form-label" for="eventName">Event Name</label>
            <input type="text" required class="form-input" id="eventName"/>

            <label class="form-label" for="description">Description</label>
            <textarea class="form-input form-textarea" id="description"></textarea>

            <label class="form-label" for="startDate">Start Date</label>
            <input type="datetime-local" required class="form-input" id="startDate"/>


            <label class="form-label" for="endDate">End Date</label>
            <input type="datetime-local" required class="form-input" id="endDate"/>

            <div class="add-button-container">
                <button class="add-button-style"
                        type="submit">
                    Add event
                </button>
            </div>
        </form>

        <div class="close-button-container">
            <button class="close-button-style" on:click={closeForm}>
                close
            </button>
        </div>
    </div>
{/if}
<div class="open-button-container">
    <button class="open-button-style" on:click={openForm}>
        +
    </button>
</div>

<style>
    .form-container {
        width: 600px;
        height: 550px;
        background-color: #fdfdfd;
        border: 2px solid #2e302e;
        border-radius: 20px;
        padding: 15px;
        position: absolute;
        margin-left: calc(50% - 300px);
        top: calc(50% - 300px);
        text-align: center;
    }

    .form-content {
        display: grid;
        text-align: left;
    }

    .form-label {
        font-size: 12pt;
        padding-top: 15px;
    }

    .form-input {
        margin-top: 5px;
        display: block;
        width: 100%;
        line-height: 24px;
        border: 1px solid rgba(0, 0, 0, 0.1);
        border-radius: 2px;
        padding-left: 5px;
        outline: none;
    }

    .form-textarea {
        height: 125px;
    }

    .add-button-container {
        position: absolute;
        bottom: 5px;
        right: 50px;
    }

    .add-button-style {
        background: #89dd39;
        color: #ffffff;
        font-size: 1.3em;
        width: 130px;
        height: 50px;
        transition: all .4s ease;
        border-radius: 5px;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .add-button-style:active {
        background: transparent;
        color: #8ee516;
    }

    .close-button-container {
        position: absolute;
        bottom: 5px;
        left: 50px;
    }

    .close-button-style {
        background: #bf3d3d;
        color: #ffffff;
        font-size: 1.3em;
        width: 100px;
        height: 50px;
        transition: all .4s ease;
        border-radius: 5px;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .close-button-style:active {
        background: transparent;
        color: #bf3d3d;
    }

    .open-button-container {
        margin-left: 50vw;
        position: fixed;
        bottom: 0;
        right: 50px;
    }

    .open-button-style {
        background: transparent;
        color: #3dbf3d;
        border: 2px solid #31de31;
        font-size: 2.5em;
        width: 60px;
        height: 60px;
        border-radius: 100%;
        transition: all .4s ease;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .open-button-style:hover {
        background: #3dbf3d;
        color: white;
    }

    .open-button-style:active {
        background: transparent;
        color: #3dbf3d;
    }
</style>
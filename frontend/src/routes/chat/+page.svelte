

<script lang='ts'>
	import InputBar from "$lib/components/messaging/input/InputBar.svelte";
	import RecyclerView from "$lib/components/messaging/recycler/RecyclerView.svelte";
    import TextSocket from "$lib/handlers/socket";
    import type {Message} from "$lib/models/chat/message/Message";
    import {size} from "$lib/stores/size";

    export let data: any;

    $: msgs = data.history.messages;

    var socket = new TextSocket((newSize: number, messages: Array<Message>) => {
        $size += newSize
        msgs = [...msgs, ...messages]
    })

</script>

<div>
    <h1 class="text-4xl mb-10 bg-primary text-black py-5 flex justify-center">
        Group Chat
    </h1>
    <RecyclerView bind:messages={msgs}/>
    <InputBar />
</div>

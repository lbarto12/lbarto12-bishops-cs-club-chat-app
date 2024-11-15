<script lang='ts'>
	import type { Message } from "$lib/models/chat/message/Message";
	import MessageComponent from "../message/Message.svelte";
    import {afterUpdate} from "svelte";


    export let messages: Array<Message>;

    $: messages;

    afterUpdate(() => {
        ref.scrollTo({top: ref.scrollHeight, behavior: 'smooth'})
    });

    let ref: HTMLDivElement;

</script>

<div bind:this={ref} class="overflow-y-scroll h-[80vh]">
    {#each messages as message}
        <MessageComponent message={message} />
    {/each}
    <div class="h-[50px]">

    </div>
</div>
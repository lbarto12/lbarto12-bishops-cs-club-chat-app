<script lang='ts'>
	import type { Message } from "$lib/models/chat/message/Message";
    import { session } from "$lib/stores/session";

    export let message: Message;
    const fromMe: boolean = $session.name == message.sender;

</script>

<div class={`relative flex pb-1 px-1 ${fromMe ? "justify-end" : "justify-start"}`}>
    <div class={`relative w-5/6 px-2 break-all rounded text-sm py-1 ${fromMe ? "bg-blue-500" : "bg-gray-200"}`}>
        {#if !fromMe}
            <div class="font-bold text-gray-600">
                {message.sender}
            </div>
        {/if}
        <span class={`${fromMe ? "text-white" : "text-gray-500"}`}>{message.content}</span>
        <span style="font-size: 10px" class="absolute mx-2 my-1 break-all text-gray-500 bottom-0 right-0">
            {message.timestamp} <!-- .toLocaleTimeString() -->
        </span>
    </div>
</div>
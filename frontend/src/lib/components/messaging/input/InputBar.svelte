

<script lang='ts'>

    import {UploadMessage} from "$lib/handlers/upload";
    import {session} from "$lib/stores/session";

    let message: string = "";

    async function submit() {
        if (message === "") return;
        await UploadMessage({
            sender: $session.name,
            timestamp: new Date(),
            content: message
        });
        message = "";
    }

</script>

<form class="fixed bottom-0 w-full" on:submit|preventDefault={submit}>
    <input bind:value={message} type="text" placeholder="Type message..." class="input input-bordered input-primary w-full">
</form>
<script lang="ts">
	import { goto } from "$app/navigation";
	import SignIn from "$lib/components/auth/SignIn.svelte";
	import { session } from "$lib/stores/session";


    let name: string;

    async function signIn() {
        $session = {
            name
        };
        
        setTimeout(() => goto('/chat'), 100);
    }

</script>

<SignIn 
    signin={signIn}
    bind:value={name}
/>

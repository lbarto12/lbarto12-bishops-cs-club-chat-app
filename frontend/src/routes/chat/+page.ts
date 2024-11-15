import { PUBLIC_API_URL } from '$env/static/public';
import {size} from "$lib/stores/size";
import type {History} from "$lib/models/chat/history/History";

/** @type {import('./$types').PageLoad} */
export async function load({params, fetch}) {
    
    const response: Response = await fetch(`${PUBLIC_API_URL}/chat`, {
        method: "GET",
    });

    const history: History = await response.json();

    size.set(history.messages.length);

    return {
        history
    }
}
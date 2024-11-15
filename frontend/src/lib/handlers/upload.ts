import type { Message } from '$lib/models/chat/message/Message';
import {PUBLIC_API_URL} from "$env/static/public";

export async function UploadMessage(message: Message) {
    const response: Response = await fetch(`${PUBLIC_API_URL}/chat`, {
        method: "POST",
        body: JSON.stringify(message)
    });

    if (response.status != 201) {
        console.log('Error creating message!');
    }

}

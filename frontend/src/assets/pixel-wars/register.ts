export async function handleSubmit(e){

    const formData = new FormData(e.target);
    console.log(formData);
    const data = formData;

    const json = {
        name: data.get('name'),
        email: data.get('email'),
        image: btoa(String.fromCharCode(...new Uint8Array(await (data.get('image') as File).arrayBuffer())))
    }

    console.log(json);
    
    try {
        await fetch('./api/profiles/new', {method: 'POST', body: JSON.stringify(json)});
    } catch (err) {
        console.log(err);
    }



}

export function getBase64(file) {
    return new Promise((resolve, reject) => {
        const fileReader = new FileReader();
        fileReader.readAsDataURL(file);

        fileReader.onload = () => {
            resolve(fileReader.result);
        };

        fileReader.onerror = (error) => {
            reject(error);
        };
    });
}
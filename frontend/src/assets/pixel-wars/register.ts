export async function handleSubmit(e){

    const formData = new FormData(e.target);

    const data = {};
    for (const field of formData) {
        const [key, value] = field;
        if (key === 'image'){
            data[key] = await getBase64(value);
            console.log(data['image']);
        }
        data[key] = value;
    }
    const json = {
        name: data['name'],
        email: data['email'],
        image: data['image']
    }

    console.log(json);
    
    try {
        await fetch('./api/profiles/new', {method: 'POST', body: json});
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
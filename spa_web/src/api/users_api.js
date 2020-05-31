const TAG_API_URL = process.env.REACT_APP_BIXPARK_CONTENT_APP_BASE_URL +
    "api/v1/users";

const UsersApi = {
    getAll: async (limit, offset) => {
        let dataUrl = TAG_API_URL + "/" + limit + "/" + offset;
        return await fetch(dataUrl, {
            method: 'GET', // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // or without this line
            redirect: 'follow'
        })
            .then((res) => res.json())
            .then((json) => json.data);
    },
    save: async (tag) => {
        let submitUrl = TAG_API_URL;
        return fetch(submitUrl, {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // or without this line
            redirect: 'follow',
            body: JSON.stringify(tag),
        }).then(res => res.json())
            .then(res => res.data);
    },
    update: async (tag) => {
        let submitUrl = TAG_API_URL;
        console.log(submitUrl, JSON.stringify(tag));
        return fetch(submitUrl, {
            method: 'PUT', // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // or without this line
            redirect: 'follow',
            body: JSON.stringify(tag),
        }).then(res => res.json())
            .then(res => res.data).catch(console.error);
    },
    delete: async (tagId) => {
        let submitUrl = TAG_API_URL + "/" + tagId;
        return fetch(submitUrl, {
            method: 'DELETE', // *GET, POST, PUT, DELETE, etc.
            mode: "cors", // or without this line
            redirect: 'follow'
        })
            .then(res => res.json())
            .then(res => res.message.message === "SUCCESS")
            .then(res => res.data).catch(err => console.log(err))
    },
};

export default UsersApi;
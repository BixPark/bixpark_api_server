import {Subject} from 'rxjs'
import UsersApi from '../api/users_api.js'

const subject = new Subject();

class UserModal {
    username = "";
    password = "";
}

const InitialData = () => {
    const initialData = new UserModal();
    return {
        username: initialData.username,
        password: initialData.password
    }
};

const initialState = {
    error: '',
    selected: InitialData(),
    authenticated: false,
    loggedUser: {}
};

let state = initialState;

const stateUpdate = state => {
    localStorage.setItem("user", state.selected);
    localStorage.setItem("authenticated", state.authenticated);
};

subject.subscribe(stateUpdate);

const authenticationStore = {
    init: async () => {
        // console.log(localStorage.getItem("authenticated"));
        state = {
            ...state,
            selected: InitialData(),
            authenticated: localStorage.getItem("authenticated") === 'true',
            loggedUser: localStorage.getItem("user")
        };
        subject.next(state);
    },
    getInitialState: InitialData,
    subscribe: setState => subject.subscribe(setState),
    authenticate: async user => {
        state = {
            ...state,
            state: 'Saving...'
        }
        subject.next(state);
        let savedData = await UsersApi.authenticate(user);

        state = {
            ...state,
            selected: InitialData(),
            loggedUser: savedData,
            authenticated: !!(savedData && savedData.data && savedData.data.ID)
        };
        subject.next(state);
    },
    initialState,

    signOut() {
        state = {
            ...state,
            selected: InitialData(),
            loggedUser: null,
            authenticated: false
        };
        subject.next(state);
    }
};

export default authenticationStore;
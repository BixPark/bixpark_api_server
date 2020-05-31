import {Subject} from 'rxjs'
import UsersApi from '../api/users_api.js'

const subject = new Subject();

class UserModal {
    first_name = "";
    last_name = "";
}

const InitialData = () => {
    const initialData = new UserModal();
    return {
        first_name: initialData.first_name,
        last_name: initialData.last_name,
    }
};

const initialState = {
    collection: [],
    error: '',
    selected: InitialData(),
    newDataCount: 0,
    limit: 10,
    offset: 0,
    status: ''
};

let state = initialState;

const userStore = {
    init: async () => {
        const tags = await UsersApi.getAll(state.limit, state.offset);
        state = {
            ...state,
            selected: InitialData(),
            collection: tags, newDataCount: tags.length
        };
        subject.next(state);
    },
    getInitialState: InitialData,
    subscribe: setState => subject.subscribe(setState),
    save: async tag => {
        state = {
            ...state,
            state: 'Saving...'
        }
        subject.next(state);
        let savedData = await UsersApi.save(tag);
        state = {
            ...state,
            collection: [...state.collection, savedData],
            newDataCount: state.newDataCount + 1
        };
        subject.next(state);
    },
    select: tag => {
        state = {
            ...state,
            selected: tag
        };
        subject.next(state);
    },
    update: async tag => {
        let updatedTag = await UsersApi.update(tag);
        state = {
            ...state,
            collection: state.collection.map((value => {
                if (value.ID === updatedTag.ID)
                    return updatedTag;
                return value;
            })),
            selected: InitialData()
        };
        subject.next(state);
    },
    initialState,
    changeCollection: async (limit, offset) => {
        state = {
            ...state,
            limit: limit,
            offset: offset,
        };
        subject.next(state);
        const tags = await UsersApi.getAll(state.limit, state.offset);
        state = {
            ...state,
            collection: tags
        }
    },
    delete: async id => {
        await UsersApi.delete(id);
        state = {
            ...state,
            collection: state.collection.filter((value => {
                return value.ID !== id;
            })),
            selected: InitialData()
        };
        subject.next(state);
    }
};

export default userStore;
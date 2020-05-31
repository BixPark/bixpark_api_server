import React, {useLayoutEffect, useState} from "react";
import userStore from '../../store/user_store'

export function List() {
    const [userState, setUserState] = useState(userStore.initialState);

    useLayoutEffect(() => {
        const subs = userStore.subscribe(setUserState);
        userStore.init().then(console.log);
        return () => subs.unsubscribe();
    }, []);


    return (
        <>

            <p className="text-xl pb-6 flex items-center">
                <i className="fas fa-list mr-3"></i> Tags
            </p>
            <div className="leading-loose">
                <table className="table-auto w-full">
                    <thead>
                    <tr>
                        <th className="px-4 py-2">ID</th>
                        <th className="px-4 py-2">Name</th>
                        <th className="px-4 py-2">Email</th>
                        <th className="px-4 py-2">Action</th>
                    </tr>
                    </thead>
                    <tbody>

                    {userState.collection.map((val, i) => {
                        return (
                            <tr className="bg-gray-100">
                                <td className="border px-4 py-2">
                                    {val.ID}
                                </td>
                                <td className="border px-4 py-2">
                                    {val.profilePic && <img
                                        href={process.env.REACT_APP_BIXPARK_CONTENT_APP_BASE_URL + "/media/" + val.profilePic.path}/>}
                                </td>
                                <td className="border px-4 py-2">
                                    <strong>{val.firstName}</strong>
                                </td>
                                <td className="border px-4 py-2">
                                    <strong>{val.lastName}</strong>
                                </td>
                                <td className="border px-4 py-2">
                                    <strong>{val.email}</strong>
                                </td>

                                <td className="border px-4 py-2">
                                    <button
                                        onClick={() =>
                                            userStore.select(val)
                                        }
                                        className="px-4  py-1 text-white  font-light tracking-wider bg-purple-500 rounded"
                                    >Edit
                                    </button>
                                    <button
                                        onClick={() =>
                                            userStore.delete(val.ID).then(
                                                () => {
                                                    userStore.changeCollection(userState.limit, userState.offset).then(console.log)
                                                }
                                            )

                                        }
                                        className="px-4  py-1 text-white  font-light tracking-wider bg-red-500 rounded"
                                    >Delete
                                    </button>
                                </td>
                            </tr>);
                    })}

                    </tbody>
                </table>
            </div>

        </>
    );
}
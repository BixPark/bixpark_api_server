import React, {useLayoutEffect, useState} from "react";
import tagStore from '../../store/user_store'

export function List() {
    const [tagSate, setTagState] = useState(tagStore.initialState);

    useLayoutEffect(() => {
        const subs = tagStore.subscribe(setTagState);
        tagStore.init().then(console.log);
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
                        <th className="px-4 py-2">Color</th>
                        <th className="px-4 py-2">Action</th>
                    </tr>
                    </thead>
                    <tbody>

                    {tagSate.collection.map((val, i) => {
                        return (
                            <tr className="bg-gray-100">
                                <td className="border px-4 py-2">
                                    {val.ID}
                                </td>
                                <td className="border px-4 py-2">
                                    <strong>{val.name}</strong>
                                </td>
                                <td className="border px-4 py-2" style={
                                    {
                                        position: "relative"
                                    }
                                }>

                                    <div style={{
                                        backgroundColor: val.primaryColor,
                                        color: val.accentColor,
                                        textAlign: "center",
                                        fontWeight: 500,
                                        width: 24,
                                        height: 24
                                    }}>
                                        P
                                    </div>
                                    <div style={{
                                        backgroundColor: val.accentColor,
                                        color: val.primaryColor,
                                        textAlign: "center",
                                        fontWeight: 500,
                                        width: 24,
                                        height: 24
                                    }}>
                                        S
                                    </div>
                                </td>
                                <td className="border px-4 py-2">
                                    <button
                                        onClick={() =>
                                            tagStore.select(val)
                                        }
                                        className="px-4  py-1 text-white  font-light tracking-wider bg-purple-500 rounded"
                                    >Edit
                                    </button>
                                    <button
                                        onClick={() =>
                                            tagStore.delete(val.ID).then(
                                                () => {
                                                    tagStore.changeCollection(tagSate.limit, tagSate.offset).then(console.log)
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
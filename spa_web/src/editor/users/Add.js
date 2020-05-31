import React, {useLayoutEffect, useState} from 'react';
import {useFormik} from "formik";
import userStore from "../../store/user_store";


function Add() {
    const [userState, setUserState] = useState(userStore.initialState);

    useLayoutEffect(() => {
        const subs = userStore.subscribe(setUserState);
        userStore.init().then(console.log);
        return () => subs.unsubscribe();
    }, []);


    const formik = useFormik({
        initialValues: userState.selected,
        enableReinitialize: true,
        onSubmit: async (values, formikBag) => {
            formikBag.setSubmitting(true);

            let data = new FormData();
            data.append("firstName", values.firstName);
            data.append("lastName", values.lastName);
            data.append("email", values.email);
            data.append("profilePic", values.profilePic);

            console.log(data);
            if (userState.selected.ID != null && userState.selected.firstName) {
                await userStore.update(data);
            } else {
                await userStore.save(data);
            }

            const newValues = {
                ...values,
                ...userStore.getInitialState()
            };
            delete newValues.ID;

            formikBag.resetForm({values: newValues});
            formikBag.setSubmitting(false);
        },
    });


    return (
        <>
            <p className="text-xl pb-6 flex items-center">
                <i className="fas fa-list mr-3"></i> Tag
                {userState.selected.ID && (<>
                    Edit
                </>)} {!userState.selected.ID && (<>
                Add
            </>)}
            </p>
            <div className="leading-loose">
                <form onSubmit={formik.handleSubmit} className="p-10 bg-white rounded shadow-xl">
                    <div className="">
                        <label className="block text-sm text-gray-600" htmlFor="firstName">First Name</label>
                        <input className="w-full px-5 py-1 text-gray-700 bg-gray-200 rounded"
                               id="firstName" name="firstName" type="text" required="" placeholder="First Name"
                               aria-label="First Name"
                               onChange={formik.handleChange}
                               value={formik.values.firstName}
                        />
                    </div>
                    <div className="">
                        <label className="block text-sm text-gray-600" htmlFor="name">Name</label>
                        <input className="w-full px-5 py-1 text-gray-700 bg-gray-200 rounded"
                               id="lastName" name="lastName" type="text" required="" placeholder="Last Name"
                               aria-label="Last Name"
                               onChange={formik.handleChange}
                               value={formik.values.lastName}
                        />
                    </div>
                    <div className="">
                        <label className="block text-sm text-gray-600" htmlFor="profilePic">Name</label>
                        <input className="w-full px-5 py-1 text-gray-700 bg-gray-200 rounded"
                               id="profilePic" name="profilePic" type="file" required="" placeholder="User Profile Pic"
                               aria-label="Profile Pic"
                               onChange={(event) => {
                                   formik.setFieldValue("profilePic", event.currentTarget.files[0]);
                               }}
                        />
                    </div>
                    <div className="mt-6">
                        <button
                            className="px-4 py-1 text-white font-light tracking-wider bg-blue-500 rounded"
                            type="submit">   {userState.status} Submit
                        </button>
                        <button
                            className="px-4 py-1 text-white font-light tracking-wider bg-red-500 rounded"
                            onClick={() => {
                                formik.resetForm();
                            }}>
                            Reset
                        </button>
                    </div>
                </form>
            </div>

        </>
    );

}

export default Add;
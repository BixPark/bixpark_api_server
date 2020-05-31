import React, {useLayoutEffect, useState} from 'react';
import {useFormik} from "formik";
import userStore from "../../store/user_store";
import {BlockPicker} from "react-color";



function AddTag() {
    const [userState, setUserState] = useState(userStore.initialState);

    useLayoutEffect(() => {
        const subs = userStore.subscribe(setUserState);
        userStore.init().then(console.log);
        return () => subs.unsubscribe();
    }, []);


    const formik = useFormik({
        initialValues: userStore.selected,
        enableReinitialize: true,
        onSubmit: async (values, formikBag) => {
            formikBag.setSubmitting(true);
            if (userStore.selected.ID != null) {
                await userStore.update(values);
            } else {
                await userStore.save(values);
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
                {userState.selected && (<>
                    Edit
                </>)} {!userState.selected && (<>
                Add
            </>)}
            </p>
            <div className="leading-loose">
                <form onSubmit={formik.handleSubmit} className="p-10 bg-white rounded shadow-xl">
                    <div className="">
                        <label className="block text-sm text-gray-600" htmlFor="name">Name</label>
                        <input className="w-full px-5 py-1 text-gray-700 bg-gray-200 rounded"
                               id="name" name="name" type="text" required="" placeholder="Your Name"
                               aria-label="Name"
                               onChange={formik.handleChange}
                               value={formik.values["name"]}
                        />
                    </div>
                    <div className="mt-2">
                        <label className="block text-sm text-gray-600" htmlFor="primaryColor">Primary
                            Color</label>
                        <BlockPicker
                            color={formik.values["primaryColor"]}
                            onChangeComplete={primaryColor => {
                                formik.setFieldValue('primaryColor', primaryColor.hex);
                            }}
                        />
                    </div>
                    <div className="mt-2">
                        <label className="block text-sm text-gray-600" htmlFor="secondaryColor">Accent
                            Color</label>
                        <BlockPicker
                            color={formik.values["accentColor"]}
                            onChangeComplete={accentColor => {
                                formik.setFieldValue('accentColor', accentColor.hex);
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

export default AddTag;
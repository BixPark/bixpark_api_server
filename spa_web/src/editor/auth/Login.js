import React, {useLayoutEffect, useState} from "react";
import "./Login.css";
import authenticationStore from "../../store/authentication_store";
import {useFormik} from "formik";

function Login() {

    const [userState, setUserState] = useState(authenticationStore.initialState);

    useLayoutEffect(() => {
        const subs = authenticationStore.subscribe(setUserState);
        authenticationStore.init().then(console.log);
        return () => subs.unsubscribe();
    }, []);


    const formik = useFormik({
        initialValues: userState.selected,
        enableReinitialize: true,
        onSubmit: async (values, formikBag) => {

            await authenticationStore.authenticate({...values});


            const newValues = {
                ...values,
                ...authenticationStore.getInitialState()
            }

            formikBag.resetForm({values: newValues});
            formikBag.setSubmitting(false);
        },
    });

    return (
        <>
            <h3 className="text-3xl">Sign In</h3>
            <form onSubmit={formik.handleSubmit} className="w-full max-w-5xl w">

                <div className="md:flex md:items-center mb-6">
                    <div className="md:w-1/3">
                        <label className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4"
                               htmlFor="inline-full-name">
                            Full Name
                        </label>
                    </div>
                    <div className="md:w-2/3">
                        <input
                            name={"username"}
                            className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
                            id="username" type="text"
                            onChange={formik.handleChange}
                            value={formik.values.username}

                        />
                    </div>
                </div>
                <div className="md:flex md:items-center mb-6">
                    <div className="md:w-1/3">
                        <label className="block text-gray-500 font-bold md:text-right mb-1 md:mb-0 pr-4"
                               htmlFor="password">
                            Password
                        </label>
                    </div>
                    <div className="md:w-2/3">
                        <input
                            className="bg-gray-200 appearance-none border-2 border-gray-200 rounded w-full py-2 px-4 text-gray-700 leading-tight focus:outline-none focus:bg-white focus:border-purple-500"
                            name={"password"}
                            id="password" type="password" placeholder="******************"
                            onChange={formik.handleChange}
                            value={formik.values.password}
                        />
                    </div>
                </div>
                <div className="md:flex md:items-center">
                    <div className="md:w-1/3"></div>
                    <div className="md:w-2/3">
                        <button
                            className="shadow bg-purple-500 hover:bg-purple-400 focus:shadow-outline focus:outline-none text-white font-bold py-2 px-4 rounded"
                            type="submit">
                            Login
                        </button>
                    </div>
                </div>
            </form>
        </>
    );
}

export default Login;

import React, {useLayoutEffect, useState} from "react";
import Login from "./editor/auth/Login";
import authenticationStore from "./store/authentication_store";
import Dashboard from "./Dashboard";

function Auth() {
    const [userState, setUserState] = useState(authenticationStore.initialState);
    useLayoutEffect(() => {
        const subs = authenticationStore.subscribe(setUserState);
        authenticationStore.init().then(console.log);
        return () => subs.unsubscribe();
    }, []);
    return (<>
        {!userState.authenticated && <div className="flex items-center justify-center h-screen w-screen">

            <div className="rounded-lg border shadow-lg p-10">
                <Login/>
            </div>
        </div>}
        {userState.authenticated && <Dashboard/>}
    </>)
}

export default Auth;

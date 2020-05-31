import React from "react";
import Add from "./Add";
import {List} from "./List";

export function UserComponent() {


    return (<>

        <h1 className="w-full text-3xl text-black pb-6">User Component</h1>

        <div className="flex flex-wrap">
            <div className="w-full lg:w-1/2 my-6 pr-0 lg:pr-2">
                <Add/>
            </div>
            <div className="w-full lg:w-1/2 mt-6 pl-0 lg:pl-2">
                <List/>
            </div>
        </div>
    </>)
}
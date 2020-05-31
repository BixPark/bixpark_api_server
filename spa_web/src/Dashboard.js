import React from 'react';
import {BrowserRouter as Router, Link, Route, Switch} from "react-router-dom";
import {UserComponent} from "./editor/users/Component";

function DashboardNav() {
    return (
        <>
            <Link to={"/"}>
                <a
                    className="flex items-center text-white opacity-75 hover:opacity-100 py-4 pl-6 nav-item">
                    <i className="fas fa-tachometer-alt mr-3"></i>
                    Dashboard
                </a>
            </Link>
            <Link to={"/users"}>
                <a
                    className="flex items-center text-white py-4 pl-6 nav-item">
                    <i className="fas fa-list mr-3"></i>
                    Users
                </a>
            </Link>
        </>
    );
}

function Dashboard() {

    return (
        <Router>

            <aside className="relative bg-sidebar h-screen w-64 hidden sm:block shadow-xl">
                <div className="p-6">
                    <a href="index.html"
                       className="text-white text-3xl font-semibold uppercase hover:text-gray-300">Admin</a>
                    <Link to={"/content"}>
                        <button
                            className="w-full bg-white cta-btn font-semibold py-2 mt-5 rounded-br-lg rounded-bl-lg rounded-tr-lg shadow-lg hover:shadow-xl hover:bg-gray-300 flex items-center justify-center">
                            <i className="fas fa-plus mr-3"></i> New Post
                        </button>
                    </Link>
                </div>
                <nav className="text-white text-base font-semibold pt-3">

                    <DashboardNav/>

                </nav>
                <a href="#"
                   className="absolute w-full upgrade-btn bottom-0 active-nav-link text-white flex items-center justify-center py-4">
                    <i className="fas fa-arrow-circle-up mr-3"></i>
                    Upgrade to Pro!
                </a>
            </aside>
            <div className="relative w-full flex flex-col h-screen overflow-y-hidden">
                <header className="w-full flex items-center bg-white py-2 px-6 hidden sm:flex">
                    <div className="w-1/2"></div>
                    <div className="relative w-1/2 flex justify-end">
                        <button
                            className="realtive z-10 w-12 h-12 rounded-full overflow-hidden border-4 border-gray-400 hover:border-gray-300 focus:border-gray-300 focus:outline-none">
                            <img src="https://source.unsplash.com/uJ8LNVCBjFQ/400x400"/>
                        </button>
                    </div>
                </header>

                <header className="w-full bg-sidebar py-5 px-6 sm:hidden">
                    <div className="flex items-center justify-between">
                        <a href="index.html"
                           className="text-white text-3xl font-semibold uppercase hover:text-gray-300">Admin</a>
                        <button className="text-white text-3xl focus:outline-none">
                            <i className="fas fa-bars"></i>
                            <i className="fas fa-times"></i>
                        </button>
                    </div>

                    <nav className="flex flex-col pt-4">
                        <DashboardNav/>
                        <button
                            className="w-full bg-white cta-btn font-semibold py-2 mt-3 rounded-lg shadow-lg hover:shadow-xl hover:bg-gray-300 flex items-center justify-center">
                            <i className="fas fa-arrow-circle-up mr-3"></i> Upgrade to Pro!
                        </button>
                    </nav>
                </header>

                <div className="w-full h-screen overflow-x-hidden border-t flex flex-col">
                    <main className="w-full flex-grow p-6">


                        <Switch>
                            <Route path="/users" component={UserComponent}/>


                            <Route path={"/"}>
                                <h1>Dashboard</h1>
                            </Route>

                        </Switch>


                    </main>

                    <footer className="w-full bg-white text-right p-4">
                        Built by <a target="_blank" href="https://ceylon.app" className="underline">Ceylon.app</a>.
                    </footer>
                </div>

            </div>

        </Router>
    );
}

export default Dashboard;
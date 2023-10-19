"use client"
import React, { useState } from 'react';
import { createUserWithEmailAndPassword } from 'firebase/auth';
import { auth } from '../../firebase/firebase';

const page = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');

  const handleSignUp = () => {
    createUserWithEmailAndPassword(auth, email, password)
      .then((userCredential) => {
     
        const user = userCredential.user;
        console.log('Signed up user:', user);
        alert('Sing-Up Successful');
      })
      .catch((error) => {
        console.error('Error signing up:', error);
      });
  };
  
  return (
    <section className="bg-gray-50 dark:bg-gray-900" style={{ minHeight: '100vh' }}>
    <div className="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0" style={{ backgroundColor: '#f4f4f4' }}>
      <a href="/" className="flex items-center mb-6 text-4xl font-semibold text-gray-900">
        HeadStart
      </a>
      <div className="w-full bg-white rounded-lg shadow md:mt-0 sm:max-w-md xl:p-0" style={{ borderColor: '#d1d5db' }}>
        <div className="p-6 space-y-4 md:space-y-6 sm:p-8">
          <h1 className="text-xl font-bold leading-tight tracking-tight text-gray-900 md:text-2xl">
            Create an account
          </h1>
          <form className="space-y-4 md:space-y-6" action="#">
            <div>
              <label htmlFor="email" className="block mb-2 text-sm font-medium text-gray-900">
                Your email
              </label>
              <input
                type="email"
                name="email"
                value = {email}
                id="email"
                onChange={(e) => setEmail(e.target.value)}
                className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                placeholder="name@company.com"

                required
              />
            </div>
            <div>
              <label htmlFor="password" className="block mb-2 text-sm font-medium text-gray-900">
                Password
              </label>
              <input
                type="password"
                name="password"
                value={password}
                id="password"
                onChange={(e) => setPassword(e.target.value)}
                placeholder="••••••••"
                className="bg-gray-50 border border-gray-300 text-gray-900 sm:text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5"
                required
              />
            </div>
            <button
              type="submit"
              onClick={handleSignUp}
              className="w-full text-black bg-gray-200 hover:bg-blue-500 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center"
              style={{ transition: 'background-color 0.3s' }}
            >
              Create an account
            </button>
            <p className="text-sm font-light text-gray-500">
              Already have an account? <a href="#" className="font-medium text-primary-600 hover:underline">Login here</a>
            </p>
          </form>
        </div>
      </div>
    </div>
  </section>
  );

}

export default page
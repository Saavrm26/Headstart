import React from "react";

const Hero = () => {
  return (
    <>
      <div className="px-4 flex flex-col items-center mt-4">
        <div className="border-b border-stroke pb-14 md:pb-20 lg:pb-24">
          <div className="-mx-4 flex flex-wrap items-center">
            <div className="w-full px-4 lg:w-1/2 2xl:w-1/2">
              <div className="max-w-[590px]">
                <h1 className="mb-6 text-3xl font-bold leading-tight text-black sm:text-4xl sm:leading-tight md:text-5xl md:leading-tight lg:text-4xl lg:leading-tight xl:text-[50px] xl:leading-tight">
                  Free Templates and Boilerplates
                </h1>
                <p className="mb-8 text-base leading-relaxed text-body-color">
                  Free, professionally designed and coded Next.js templates and
                  boilerplates that can be used for your next web projects.
                  These pre-made React-Next.js templates and boilerplates are
                  easy to configure and customize so you can quickly get your
                  project up and running without designing from scratch or
                  coding the frontend.
                </p>
              </div>
            </div>
            <div className="w-full px-4 lg:w-1/2 2xl:w-1/2">
              <div className="relative z-10 mx-auto mt-12 w-full text-center lg:mt-0 lg:text-right">
                <span
                  style={{
                    boxSizing: "border-box",
                    display: "inline-block",
                    overflow: "hidden",
                    width: "initial",
                    height: "initial",
                    background: "none",
                    opacity: 1,
                    border: "0px",
                    margin: "0px",
                    padding: "0px",
                    position: "relative",
                    maxWidth: "100%",
                  }}
                >
                  <span
                    style={{
                      boxSizing: "border-box",
                      display: "block",
                      width: "initial",
                      height: "initial",
                      background: "none",
                      opacity: 1,
                      border: "0px",
                      margin: "0px",
                      padding: "0px",
                      maxWidth: "100%",
                    }}
                  >
                    <img
                      alt=""
                      aria-hidden="true"
                      src="data:image/svg+xml,%3csvg%20xmlns=%27http://www.w3.org/2000/svg%27%20version=%271.1%27%20width=%27454%27%20height=%27315%27/%3e"
                      style={{
                        display: "block",
                        maxWidth: "100%",
                        width: "initial",
                        height: "initial",
                        background: "none",
                        opacity: 1,
                        border: "0px",
                        margin: "0px",
                        padding: "0px",
                      }}
                    />
                  </span>
                  <img
                    alt="hero image"
                    srcSet="https://nextjstemplates.com/_next/static/media/hero-image.9facdb18.svg"
                    src="/_next/static/media/hero-image.9facdb18.svg"
                    decoding="async"
                    data-nimg="intrinsic"
                    className="mx-auto lg:ml-auto lg:mr-0"
                    style={{
                      position: "absolute",
                      inset: 0,
                      boxSizing: "border-box",
                      padding: 0,
                      border: "none",
                      margin: "auto",
                      display: "block",
                      width: 0,
                      height: 0,
                      minWidth: "100%",
                      maxWidth: "100%",
                      minHeight: "100%",
                      maxHeight: "100%",
                    }}
                  />
                </span>
              </div>
            </div>
          </div>
          <button className="bg-blue-700 text-white p-2 lg:p-3 border rounded-lg mx-[50%] lg:mx-2">All Templates</button>
        </div>
      </div>
    </>
  );
};

export default Hero;

// 04832230

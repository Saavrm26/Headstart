import React from "react";

async function getData() {
  const res = await fetch('http://localhost:8080/catalouge/')
  if (!res.ok) {
    throw new Error('Failed to fetch data')
  }
 
  return res.json()
}

type Template = {
  description: string;
  downloadURL: string;
  imageURL: string;
  name: string;
};


const Templates = async() => {
  const data = await getData()
  console.log(data);
  return (
    <>
      <section id="featured templates" className=" py-14 md:py-20 lg:py-24">
        <div className="w-full">
          <div className="mx-4 flex flex-wrap">
            <div className="w-full px-4">
              <div className="mx-auto mb-12 max-w-[690px] text-center">
                <h2 className="mb-3 text-center text-3xl font-bold !leading-tight text-black md:text-4xl xl:text-[45px]">
                  Sample Templates
                </h2>
                <p className="text-base leading-relaxed text-body-color">
                  Standard Templates that will make your life easy
                </p>
              </div>
            </div>
          </div>

          {/* featured templates */}
          <div className="flex items-center justify-center">
            <div className="grid grid-cols-1 gap-8 md:grid-cols-2">
              {/* standard templates */}
              {
                data.map((template:Template,index:any)=>(
                <div key={index} className="relative flex flex-col justify-between rounded-md border border-stroke bg-white shadow-one transition-transform duration-500 hover:scale-105">
                <div className="p-4">
                  <div className="relative aspect-[81/41] w-full overflow-hidden rounded border border-stroke">
                    <span
                      style={{
                        position: "absolute",
                        inset: 0,
                        boxSizing: "border-box",
                        padding: 0,
                        border: "none",
                        margin: 0,
                        display: "block",
                        width: "initial",
                        height: "initial",
                        background: "none",
                        opacity: 1,
                      }}
                    >
                      <img srcSet="https://nextjstemplates.com/_next/image?url=https%3A%2F%2Fapi.nextjstemplates.com%2Fimage%2FAI-Tool---OpenAI-Next.js-SaaS-Starter-Kit-443b781b-bcb8-4198-9e19-56232a3484d6.png&w=1920&q=100" />
                    </span>
                  </div>
                </div>
                <div className="flex flex-1 flex-col justify-between">
                  <div className="flex-1 px-6 pt-2">
                    <h3 className="mb-4 text-xl font-semibold text-[#333333] duration-300 hover:text-primary">
                      <a className="block" href="/templates/ai-tool">
                        AI Tool - OpenAI Next.js SaaS Starter Kit
                      </a>
                    </h3>
                    <p className="mb-5 text-sm leading-relaxed text-body-color">
                      {template.description}
                    </p>
                  </div>
                  <div className="relative flex flex-wrap justify-around border-t border-stroke after:absolute after:left-1/2 after:top-0 after:h-full after:w-[1px] after:-translate-x-1/2 after:bg-stroke">
                    <a className="flex items-center justify-center py-3 text-sm font-medium text-body-color hover:text-primary">
                      Preview
                    </a>
                    <div className="border border-b-slate-300 w-0"></div>
                    <a className="flex items-center justify-center py-3 text-sm font-medium text-body-color hover:text-primary">
                      Download
                    </a>
                  </div>
                </div>
              </div>
                ))
              }
            </div>
          </div>
        </div>
      </section>
    </>
  );
};

export default Templates;

import React from "react";
import ScrollUp from "@/components/Common/ScrollUp";
import Footer from "@/components/Footer";
import Header from "@/components/Header";

const Page = () => {
  return (
    <div className="flex flex-col min-h-screen">
      <ScrollUp />
      <Header />
      <div className="flex-1">
        {/* Content of your page goes here */}
      </div>
      <Footer />
    </div>
  );
};

export default Page;

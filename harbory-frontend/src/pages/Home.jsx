import { AppWindow, Cuboid, Image, Inbox, Wifi } from 'lucide-react'
import React from 'react'

const Home = () => {
  return (
    <div className='h-screen w-screen bg-gray-900 overflow-hidden m-0 flex'>
      <div className='w-64 bg-gray-800 p-4'>
        <nav className='space-y-2'>
          <div className='flex items-center space-x-3 bg-gray-700 p-3 rounded-lg text-white cursor-pointer'>
            {/* <div className='w-5 h-5 bg-blue-500 rounded'></div> */}
            <AppWindow className='w-5 h-5' />
            <span>Dashboard</span>
          </div>
          <div className='flex items-center space-x-3 p-3 rounded-lg text-gray-300 hover:bg-gray-700 cursor-pointer' onClick={() => window.location.href = '/containers'}>
            {/* <div className='w-5 h-5 border-2 border-gray-400 rounded'></div> */}
            <Cuboid className='w-5 h-5' />
            <span>Containers</span>
          </div>
          <div className='flex items-center space-x-3 p-3 rounded-lg text-gray-300 hover:bg-gray-700 cursor-pointer'>
            {/* <div className='w-5 h-5 border-2 border-gray-400 rounded'></div> */}
            <Image className='w-5 h-5' />
            <span>Images</span>
          </div>
          <div className='flex items-center space-x-3 p-3 rounded-lg text-gray-300 hover:bg-gray-700 cursor-pointer'>
            {/* <div className='w-5 h-5 border-2 border-gray-400 rounded'></div> */}
            <Inbox className='w-5 h-5' />
            <span>Volumes</span>
          </div>
          <div className='flex items-center space-x-3 p-3 rounded-lg text-gray-300 hover:bg-gray-700 cursor-pointer'>
            {/* <div className='w-5 h-5 border-2 border-gray-400 rounded'></div> */}
            <Wifi className='w-5 h-5' />
            <span>Networks</span>
          </div>
        </nav>
      </div>
      
      <div className='flex-1 p-8'>
        <h1 className='text-4xl font-bold text-white mb-8'>Dashboard</h1>
        
        <div className='mb-8'>
          <h2 className='text-2xl font-semibold text-white mb-6'>System Overview</h2>
          
          <div className='grid grid-cols-2 gap-8'>
            <div className='space-y-4'>
              <div className='flex justify-between items-center border-b border-gray-700 pb-3'>
                <span className='text-gray-300'>Docker Version</span>
                <span className='text-white'>24.0.5</span>
              </div>
              <div className='flex justify-between items-center border-b border-gray-700 pb-3'>
                <span className='text-gray-300'>Images</span>
                <span className='text-white'>10</span>
              </div>
              <div className='flex justify-between items-center border-b border-gray-700 pb-3'>
                <span className='text-gray-300'>Volumes</span>
                <span className='text-white'>1</span>
              </div>
              <div className='flex justify-between items-center border-b border-gray-700 pb-3'>
                <span className='text-gray-300'>Operating System</span>
                <span className='text-white'>Ubuntu 22.04.3 LTS</span>
              </div>
            </div>
            
            <div className='space-y-4'>
              <div className='flex justify-between items-center border-b border-gray-700 pb-3'>
                <span className='text-gray-300'>Containers</span>
                <span className='text-white'>3</span>
              </div>
              <div className='flex justify-between items-center border-b border-gray-700 pb-3'>
                <span className='text-gray-300'>Networks</span>
                <span className='text-white'>2</span>
              </div>
              <div className='flex justify-between items-center border-b border-gray-700 pb-3'>
                <span className='text-gray-300'>Uptime</span>
                <span className='text-white'>2 days 14 hours</span>
              </div>
              <div className='flex justify-between items-center border-b border-gray-700 pb-3'>
                <span className='text-gray-300'>Architecture</span>
                <span className='text-white'>x86_64</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Home

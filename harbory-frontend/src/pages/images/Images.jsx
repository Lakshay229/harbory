import React, { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Navbar from '../../components/Navbar'

const Images = () => {
    const [images, setImages] = useState([])
    const [loading, setLoading] = useState(false)
    const [error, setError] = useState(null)
    const navigate = useNavigate()

    const apiurl = import.meta.env.VITE_API_URL
    
      const fetchImages = async () => {
        try {
          setLoading(true)
          const response = await fetch(`${apiurl}/api/images`)
          if (!response.ok) {
            throw new Error('Failed to fetch images')
          }
          const data = await response.json()
          setImages(data)
          setError(null)
        } catch (err) {
          setError(err.message)
        } finally {
          setLoading(false)
        }
      }
    
      useEffect(() => {
        fetchImages()
      }, [])
      
      const formatCreated = (timestamp) => {
    return new Date(timestamp * 1000).toLocaleDateString()
  }

  return (
    <div className='bg-gray-950 w-screen'>
      <Navbar />
      <div className='flex-1 overflow-auto p-4'>
          <div className='flex justify-between items-center mb-6'>
            <h1 className='text-white text-3xl font-bold'>
              Images
            </h1>
            <div className='flex gap-2'>
              <button 
                className='bg-blue-600 hover:bg-blue-700 text-white px-2 py-2 rounded text-sm'
                onClick={fetchImages}
              >
                Refresh
              </button>
            </div>
          </div>

          {loading ? (
            <div className='flex justify-center items-center h-64'>
              <div className='text-white text-lg'>Loading images...</div>
            </div>
          ) : error ? (
            <div className='bg-red-900 border border-red-700 rounded-lg p-4 text-red-200'>
              <h3 className='font-bold mb-2'>Error</h3>
              <p>{error}</p>
              <button 
                onClick={fetchImages}
                className='mt-2 bg-red-700 hover:bg-red-600 text-white px-2 py-2 rounded text-sm'
              >
                Retry
              </button>
            </div>
          ) : (
            <div className='bg-gray-800 rounded-lg overflow-hidden shadow-lg'>
                <table className='table-fixed w-full text-sm text-left'>
                  <thead className='text-xs text-gray-300 uppercase bg-gray-700'>
                    <tr>
                      <th className='px-4 py-3'>Project</th>
                      <th className='px-4 py-3'>Service</th>
                      <th className='px-4 py-3'>Tag</th>
                      <th className='px-4 py-3'>Size</th>
                      <th className='px-4 py-3'>Created</th>
                      <th className='px-4 py-3'>Actions</th>
                    </tr>
                  </thead>
                  <tbody>
                    {images.map((image) => (
                      <tr key={image.Id} className='bg-gray-800 border-b border-gray-700 hover:bg-gray-750'>
                        <td className='px-4 py-4'>
                          <div className='font-medium text-white'>
                            {image.Labels['com.docker.compose.project'] || 'N/A'}
                          </div>
                        </td>
                        <td className='px-4 py-4 text-gray-300'>
                          {image.Labels['com.docker.compose.service'] || 'N/A'}
                        </td>
                        <td className='px-4 py-4 text-gray-300'>
                          <div className='max-w-xs truncate' title={image.RepoTags.length === 0 ? 'none' : image.RepoTags.join(', ')}>
                            {image.RepoTags.length === 0 ? 'none' : image.RepoTags.join(', ')}
                          </div>
                        </td>
                        <td className='px-4 py-4 text-gray-300'>
                          {(image.Size / (1024 * 1024)).toFixed(2)} MB
                        </td>
                        <td className='px-4 py-4 text-gray-300'>
                          {formatCreated(image.Created)}
                        </td>
                        <td className='px-4 py-4'>
                          <button 
                            className='text-blue-400 hover:text-blue-300 font-medium transition-colors'
                            onClick={() => navigate(`/images/${image.Id}`)}
                          >
                            Inspect
                          </button>
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              
            </div>
          )}

          {!loading && !error && images.length === 0 && (
            <div className='bg-gray-800 rounded-lg p-8 text-center'>
              <div className='text-gray-400 text-lg mb-2'>No images found</div>
              <div className='text-gray-500 text-sm'>No Docker images are currently available</div>
            </div>
          )}
        </div>
    </div>
  )
}

export default Images

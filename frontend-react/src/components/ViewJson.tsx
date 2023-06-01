import ReactJson from 'react-json-view'

const ViewJson = ({ json }: { json: object }) => {
  return (
    <ReactJson
      src={json}
      name={false}
      theme="colors"
      style={{ textAlign: 'left' }}
    />
  )
}

export default ViewJson
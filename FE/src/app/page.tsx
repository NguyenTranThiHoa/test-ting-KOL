"use client"

import React, { useEffect, useRef, useState } from "react"
import axios from "axios"
import "./page.css"

interface Kols {
	kolID: number;
	code: string;
	portraitURL: string;
	language: string;
	expectedSalary: number;
  }

const Page = () => {
  const [kols, setKols] = useState<Kols[]>([])
  const scrollRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    axios.get("http://localhost:8081/kols")
      .then(res => {
        setKols(res.data.kol)
      })
      .catch(err => console.error(err))
  }, [])

  const scrollLeft = () => {
    if (scrollRef.current) {
      scrollRef.current.scrollBy({ left: -300, behavior: 'smooth' })
    }
  }

  const scrollRight = () => {
    if (scrollRef.current) {
      scrollRef.current.scrollBy({ left: 300, behavior: 'smooth' })
    }
  }

  return (
    <>
      <h1 className="header">Danh sách KOL</h1>

      <div className="scroll-buttons">
        <button onClick={scrollLeft}>⬅️</button>
        <button onClick={scrollRight}>➡️</button>
      </div>

      <div className="kol-list" ref={scrollRef}>
	  { kols.map(kol => (
		<div key={kol.kolID} className="kol-card">
			<img src={kol.portraitURL} alt={`Portrait of ${kol.code}`} />
			<p><strong>Code:</strong> {kol.code}</p>
			<p><strong>Language:</strong> {kol.language}</p>
			<p><strong>Salary:</strong> ${kol.expectedSalary}</p>
		</div>
		))}
      </div>
    </>
	)
	
	
}

export default Page

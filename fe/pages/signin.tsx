import { Metadata } from "next"
import Image from "next/image"
import Link from "next/link"
import { Command } from "lucide-react"

import { cn } from "@/lib/utils"
import { buttonVariants } from "@/components/button"
import SignUpForm from "@/components/user-auth-form"
import { Signin } from '@/components/signin'

export default function Signup() {
    return (
        <>
            <div className="container relative h-screen flex-col items-center justify-center lg:px-0">
                <Link
                    href="/signup"
                    className={cn(
                        buttonVariants({ variant: "ghost" }),
                        "absolute right-0 top-0 mr-2"
                    )}
                >
                    SignUp
                </Link>
                <div className="relative z-20 flex items-center text-lg font-medium m-6">
                    <Command className="mr-2 h-6 w-6" /> BolusWizard.io
                </div>
                <div className="container max-w-md">
                    <Signin />
                </div>
            </div>
        </>
    )
}
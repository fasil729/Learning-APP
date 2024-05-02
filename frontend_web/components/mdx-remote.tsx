// components/mdx-remote.js
import { MDXRemote } from 'next-mdx-remote/rsc'
import { Test } from './test'
import { twMerge } from "tailwind-merge"
import { ClassValue, clsx } from "clsx"
import Link from 'next/link'
import ImagePopup from './imagePopUp'

interface CalloutProps {
    icon?: string
    children?: React.ReactNode
    type?: "default" | "warning" | "danger"
  }
  
export function Callout({
    children,
    icon,
    type = "default",
    ...props
  }: CalloutProps) {
    return (
      <div
        className={cn("my-6 flex items-start rounded-md border border-l-4 p-4", {
          "border-red-900 bg-red-50": type === "danger",
          "border-yellow-900 bg-yellow-50": type === "warning",
        })}
        {...props}
      >
        {icon && <span className=" mr-4 text-2xl">{icon}</span>}
        <div>{children}</div>
      </div>
    )
  }
export function cn(...inputs: ClassValue[]) {
    return twMerge(clsx(inputs))
  }


  interface CardProps extends React.HTMLAttributes<HTMLDivElement> {
    href?: string
    disabled?: boolean
  }
  
  export function MdxCard({
    href,
    className,
    children,
    disabled,
    ...props
  }: CardProps) {
    return (
      <div
        className={cn(
          "group relative rounded-lg border p-6 shadow-md transition-shadow hover:shadow-lg",
          disabled && "cursor-not-allowed opacity-60",
          className
        )}
        {...props}
      >
        <div className="flex flex-col justify-between space-y-4">
          <div className="space-y-2 [&>h3]:!mt-0 [&>h4]:!mt-0 [&>p]:text-muted-foreground">
            {children}
          </div>
        </div>
        {href && (
          <Link href={disabled ? "#" : href} className="absolute inset-0">
            <span className="sr-only">View</span>
          </Link>
        )}
      </div>
    )
  }

const components = {
    h1: ({ className, ...props }) => (
      <h1
        className={cn(
          "mt-2 scroll-m-20 text-4xl font-bold tracking-tight",
          className
        )}
        {...props}
      />
    ),
    h2: ({ className, ...props }) => (
      <h2
        className={cn(
          "mt-10 scroll-m-20 border-b pb-1 text-3xl font-semibold tracking-tight first:mt-0",
          className
        )}
        {...props}
      />
    ),
    h3: ({ className, ...props }) => (
      <h3
        className={cn(
          "mt-8 scroll-m-20 text-2xl font-semibold tracking-tight",
          className
        )}
        {...props}
      />
    ),
    h4: ({ className, ...props }) => (
      <h4
        className={cn(
          "mt-8 scroll-m-20 text-xl font-semibold tracking-tight",
          className
        )}
        {...props}
      />
    ),
    h5: ({ className, ...props }) => (
      <h5
        className={cn(
          "mt-8 scroll-m-20 text-lg font-semibold tracking-tight",
          className
        )}
        {...props}
      />
    ),
    h6: ({ className, ...props }) => (
      <h6
        className={cn(
          "mt-8 scroll-m-20 text-base font-semibold tracking-tight",
          className
        )}
        {...props}
      />
    ),
    a: ({ className, ...props }) => (
      <a
        className={cn("font-medium underline underline-offset-4", className)}
        {...props}
      />
    ),
    p: ({ className, ...props }) => (
      <p
        className={cn("leading-7 [&:not(:first-child)]:mt-6", className)}
        {...props}
      />
    ),
    ul: ({ className, ...props }) => (
      <ul className={cn("my-6 ml-6 list-disc", className)} {...props} />
    ),
    ol: ({ className, ...props }) => (
      <ol className={cn("my-6 ml-6 list-decimal", className)} {...props} />
    ),
    li: ({ className, ...props }) => (
      <li className={cn("mt-2", className)} {...props} />
    ),
    blockquote: ({ className, ...props }) => (
      <blockquote
        className={cn(
          "mt-6 border-l-2 pl-6 italic [&>*]:text-muted-foreground",
          className
        )}
        {...props}
      />
    ),
    img: ({
      className,
      alt,
      ...props
    }: React.ImgHTMLAttributes<HTMLImageElement>) => (
        
      // eslint-disable-next-line @next/next/no-img-element
      <img className={cn("rounded-md border", className)} alt={alt} {...props} />
    ),
    hr: ({ ...props }) => <hr className="my-4 md:my-8" {...props} />,
    table: ({ className, ...props }: React.HTMLAttributes<HTMLTableElement>) => (
      <div className="my-6 w-full overflow-y-auto">
        <table className={cn("w-full", className)} {...props} />
      </div>
    ),
    tr: ({ className, ...props }: React.HTMLAttributes<HTMLTableRowElement>) => (
      <tr
        className={cn("m-0 border-t p-0 even:bg-muted", className)}
        {...props}
      />
    ),
    th: ({ className, ...props }) => (
      <th
        className={cn(
          "border px-4 py-2 text-left font-bold [&[align=center]]:text-center [&[align=right]]:text-right",
          className
        )}
        {...props}
      />
    ),
    td: ({ className, ...props }) => (
      <td
        className={cn(
          "border px-4 py-2 text-left [&[align=center]]:text-center [&[align=right]]:text-right",
          className
        )}
        {...props}
      />
    ),
    pre: ({ className, ...props }) => (
      <pre
        className={cn(
          "mb-4 mt-6 overflow-x-auto rounded-lg border bg-black py-4",
          className
        )}
        {...props}
      />
    ),
    code: ({ className, ...props }) => (
      <code
        className={cn(
          "relative rounded px-[0.3rem] py-[0.2rem] font-mono text-sm text-white",
          className
        )}
        {...props}
      />
    ),
    Callout,
    MdxCard,
    Image: ImagePopup,
    ImagePlaceholder: ImagePopup,

}

export function CustomMDX(props) {
  return (
    <MDXRemote
      {...props}
      components={{ ...components, ...(props.components || {}) }}
    />
  )
}


import { CustomMDX } from "@/components/mdx-remote"

interface Props {
     data: any
}

export const MdxComponent = ({data}: Props) => {
    return <CustomMDX
        source={data}
        />
}
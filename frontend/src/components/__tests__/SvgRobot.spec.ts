import { describe, it, expect } from 'vitest'
import { mount } from '@vue/test-utils'
import { ref } from 'vue'
import SvgRobot from '../SvgRobot.vue'
import SvgText from '../SvgText.vue'

const globalProvide = {
  global: {
    provide: {
      'rotate-field': ref(false),
    },
  },
}

describe('SvgRobot', () => {
  const defaultProps = {
    x: 1,
    y: 2,
    orientation: 0,
    id: 5,
    teamColor: 'YELLOW' as const,
  }

  it('renders a path and SvgText', () => {
    const wrapper = mount(SvgRobot, {
      props: defaultProps,
      ...globalProvide,
    })
    expect(wrapper.find('path').exists()).toBe(true)
    expect(wrapper.findComponent(SvgText).exists()).toBe(true)
  })

  it('passes robot id as text to SvgText', () => {
    const wrapper = mount(SvgRobot, {
      props: { ...defaultProps, id: 7 },
      ...globalProvide,
    })
    expect(wrapper.findComponent(SvgText).props('text')).toBe('7')
  })

  it('uses yellow fill for YELLOW team', () => {
    const wrapper = mount(SvgRobot, {
      props: { ...defaultProps, teamColor: 'YELLOW' },
      ...globalProvide,
    })
    const style = wrapper.find('path').attributes('style')
    expect(style).toContain('fill: yellow')
  })

  it('uses blue fill for BLUE team', () => {
    const wrapper = mount(SvgRobot, {
      props: { ...defaultProps, teamColor: 'BLUE' },
      ...globalProvide,
    })
    const style = wrapper.find('path').attributes('style')
    expect(style).toContain('fill: blue')
  })

  it('generates a valid SVG path with arc command', () => {
    const wrapper = mount(SvgRobot, {
      props: defaultProps,
      ...globalProvide,
    })
    const d = wrapper.find('path').attributes('d')
    expect(d).toMatch(/^M .+ A .+ L .+$/)
  })
})
